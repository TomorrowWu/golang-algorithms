package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	// 可以把我们的区块链用漂亮的格式打印到终端terminal中
	"github.com/joho/godotenv" // 允许我们从之前创建的.evn文件读取配置
)

//Block 区块定义
type Block struct {
	Index     int // 区块高度
	Timestamp string
	BPM       int // 脉搏
	Hash      string
	PrevHash  string
	Validator string
}

//Blockchain 存储区块链
var Blockchain []Block
var tempBlocks []Block // 临时存储单元，在区块被选出来并添加到BlockChain之前，临时存储在这里

var candidateBlocks = make(chan Block) // candidateBlocks是Block的通道，任何一个节点在提出一个新块时都将它发送到这个通道

var announcements = make(chan string) // announcements也是一个通道，我们的主Go TCP服务器将向所有节点广播最新的区块链

var mutex = &sync.Mutex{} // mutex是一个标准变量，允许我们控制读/写和防止数据竞争

var validators = make(map[string]int) // validators是节点的存储map，同时也会保存每个节点持有的令牌数

func main() {
	err := godotenv.Load() // 读取环境变量
	if err != nil {
		log.Fatal(err)
	}

	// 生成创世区块
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, calculateBlockHash(genesisBlock), "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	httpPort := os.Getenv("PORT") // 设置监听的端口

	// 启动一个tcp服务器
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :", httpPort)
	defer server.Close()

	// 将候选区块加入到tempBlocks中
	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	// 选出胜利者
	go func() {
		for {
			pickWinner()
		}
	}()

	// 不断的接收请求，处理请求
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

// pickWinner creates a lottery pool of validators and chooses the validator who gets to forge a block to the blockchain
// by random selecting from the pool, weighted by amount of tokens staked
func pickWinner() {
	time.Sleep(30 * time.Second) // 每三十秒执行一次
	mutex.Lock()
	temp := tempBlocks
	mutex.Unlock()

	lotteryPool := []string{} // 乐透池
	// 对于提议块的暂存区域，我们会通过if len(temp) > 0来判断是否已经有了被提议的区块。
	if len(temp) > 0 {

		// 稍微针对传统的proof of stake算法修改了一下
		// 所有的节点都需要提交一个区块，区块的权重由节点下注的token数量决定
		// 传统的proof of stake算法，验证者节点无需铸造新的区块也能获得收益
	OUTER:
		for _, block := range temp {
			// 如果已经在乐透池中，忽略节点
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			// 锁定节点列表，防止数据静态条件
			mutex.Lock()
			setValidators := validators
			mutex.Unlock()

			// 那么我们怎么根据这些验证者持有的令牌数来给予他们合适的随机权重呢？
			// 首先，用验证者的令牌填充lotteryPool数组，例如一个验证者有100个令牌，那么在lotteryPool中就将有100个元素填充；
			// 如果有1个令牌，那么将仅填充1个元素。
			k, ok := setValidators[block.Validator] // k为验证者节点的token数量
			if ok {
				for i := 0; i < k; i++ {
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}
		}

		// 从乐透池中随机选择一个胜者
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		// 将胜者提交的区块添加到区块链中，然后让其他所有节点都知道
		for _, block := range temp {
			if block.Validator == lotteryWinner {
				mutex.Lock()
				Blockchain = append(Blockchain, block)
				mutex.Unlock()
				for range validators {
					announcements <- "\nwinning validator: " + lotteryWinner + "\n"
				}
				break
			}
		}
	}

	mutex.Lock()
	tempBlocks = []Block{}
	mutex.Unlock()
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 第一个Go协程接收并打印出来自TCP服务器的任何通知，这些通知包含了获胜验证者的通知。
	go func() {
		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()
	// 验证者节点的地址
	var address string

	// 允许用户可以分配一定数量的token来下注
	// token数量越多，就有越大的概率获得新的区块的奖励
	io.WriteString(conn, "Enter token balance:")
	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text()) // 字符串转为整数
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
			return
		}
		t := time.Now()
		address = calculateHash(t.String())
		validators[address] = balance
		fmt.Println(validators)
		break
	}

	io.WriteString(conn, "\nEnter a new BPM:") // 输入要上链的心率

	scanBPM := bufio.NewScanner(conn)

	go func() {
		for {
			// 从stdin获取心率值，然后完成一些校验并将其加入到区块链上
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())
				// 如果有恶意节点想要使用一个有问题的输入来破坏区块链，将这个恶意节点从validators中删除掉, 他们也将丢失他们持有的token。
				if err != nil {
					log.Printf("%v not a number: %v", scanBPM.Text(), err)
					delete(validators, address)
					conn.Close()
				}

				mutex.Lock()
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()

				// 铸造一个候选区块
				newBlock, err := generateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if isBlockValid(newBlock, oldLastIndex) {
					candidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nEnter a new BPM:")
			}
		}
	}()

	// 模拟接收到广播
	// 最后一段的循环会周期性的打印出最新的区块链，这样每个验证者都能获知最新的状态
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		output, err := json.Marshal(Blockchain)
		mutex.Unlock()
		if err != nil {
			log.Fatal(err)
		}
		// 写入conn中，客户端会打印
		io.WriteString(conn, string(output)+"\n")
	}

}

// 校验区块合法性
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// 计算哈希的函数
func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 就是简单的计算一个哈希出来，并没有挖矿。
func calculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	return calculateHash(record)
}

// 产生新的区块，需要使用上一个区块的哈希
func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil
}
