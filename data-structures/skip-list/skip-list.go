package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

//TODO 吴名 2018/11/19 11:36 implement is so complex

const (
	// MAX_LEVEL max level of skip-list
	MAX_LEVEL = 16
)

// Node defines a node struct of skip-list
type Node struct {
	v        interface{} //value
	score    int         //score used for sorting
	level    int         //height
	forwards []*Node     //forward pointer per layer
}

func newSkipListNode(v interface{}, score, level int) *Node {
	return &Node{
		v:        v,
		score:    score,
		level:    level,
		forwards: make([]*Node, level, level),
	}
}

// SkipList defines a skip-list
type SkipList struct {
	head   *Node
	level  int
	length int
}

// NewSkipList returns a skip-list object
func NewSkipList() *SkipList {
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

// Length returns the length of skip-list
func (sl *SkipList) Length() int {
	return sl.length
}

// Level returns the level of skip-list
func (sl *SkipList) Level() int {
	return sl.level
}

// Insert inserts a node into skip-list
func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	cur := sl.head
	update := [MAX_LEVEL]*Node{}
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	newNode := newSkipListNode(v, score, level)

	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	if level > sl.level {
		sl.level = level
	}

	sl.length++

	return 0
}

// Find 查找
func (sl *SkipList) Find(v interface{}, score int) *Node {
	if nil == v || sl.length == 0 {
		return nil
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

// Delete 删除节点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找前驱节点
	cur := sl.head
	//记录前驱路径
	update := [MAX_LEVEL]*Node{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--

	return 0
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}
