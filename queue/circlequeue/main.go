package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int //4
	array   [5]int
	head    int //指向队列队首 0,包含元素
	tail    int //指向队尾 0,不含最后元素
}

func (circleQueue *CircleQueue) Push(val int) (err error) {
	if circleQueue.IsFull() {
		return errors.New("queue is full")
	}
	//queue.tail在队列尾部,不包含最后的元素
	circleQueue.array[circleQueue.tail] = val //把值给尾部
	circleQueue.tail = (circleQueue.tail + 1) % circleQueue.maxSize
	return
}

func (circleQueue *CircleQueue) Pop() (val int, err error) {
	if circleQueue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	val = circleQueue.array[circleQueue.head]
	circleQueue.head = (circleQueue.head + 1) % circleQueue.maxSize

	return val, nil
}

func (circleQueue *CircleQueue) ListQueue() {
	//取出当前队列有多少个元素
	size := circleQueue.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}

	//设计一个辅助的变量,指向head
	tempHead := circleQueue.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, circleQueue.array[tempHead])
		tempHead = (tempHead + 1) % circleQueue.maxSize
	}
	fmt.Println()
}

func (circleQueue *CircleQueue) IsFull() bool {
	//会预留一个空位,
	return (circleQueue.tail+1)%circleQueue.maxSize == circleQueue.head
}

func (circleQueue *CircleQueue) IsEmpty() bool {
	return circleQueue.tail == circleQueue.head
}

//取出环形队列有多少个元素
func (circleQueue *CircleQueue) Size() int {
	//这是一个关键的算法.
	//演算一下,size不可能超过maxSize,存在tail<head的情况,所以直接+maxSize,再取模
	return (circleQueue.tail + circleQueue.maxSize - circleQueue.head) % circleQueue.maxSize
}

func main() {
	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
