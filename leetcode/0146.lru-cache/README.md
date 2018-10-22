### [146. LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/description/)
运用你所掌握的数据结构，设计和实现一个  **LRU (最近最少使用)** 缓存机制。它应该支持以下操作： 获取数据 **get** 和 写入数据 **put** 。

获取数据 **get(key)** - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 **put(key, value)** - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

**进阶:**

你是否可以在 **O(1)** 时间复杂度内完成这两种操作？

**示例:**
```
LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
```

### 解题思路
- LRU是Least Recently Used的缩写,即"最近最少使用",也就是说,LRU缓存把最近最少使用的数据移除,让给最新读取的数据
- 往往最常读取的,也是读取次数最多的
- 操作系统等最常使用的缓存策略,即LRU
- 需要在O(1)时间复杂度实现put操作,就需要使用双向链表,方便移动节点(单链表无法达到O(1))
- O(1)实现get查询操作,就需要使用map存key-node键值对,实现快速查询
- 具体详见代码注释

### 代码实现

```Golang
// doublyLinkedNode defines a node for double-linked-list
type doublyLinkedNode struct {
	prev, next *doublyLinkedNode
	key, val   int
}

// LRUCache defines a object for cache
type LRUCache struct {
	len, cap    int
	first, last *doublyLinkedNode         //head,tail
	nodes       map[int]*doublyLinkedNode //hashtable,find node in O(1)
}

// Constructor creates a cache object
func Constructor(capacity int) LRUCache {
	return LRUCache{
		len:   0,
		cap:   capacity,
		first: nil,
		last:  nil,
		nodes: make(map[int]*doublyLinkedNode, capacity),
	}
}

// Get returns value by key
func (l *LRUCache) Get(key int) int {
	if node, ok := l.nodes[key]; ok {
		//key exist
		// move the node to the head of double-linked-list
		l.moveToFirst(node)
		return node.val
	}

	//key not exist,return -1
	return -1
}

// Put puts key-value pair into LRUCache
func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.nodes[key]; ok {
		//update value of old node
		node.val = value
		// move the node to the head of double-linked-list
		l.moveToFirst(node)
	} else {
		if l.len == l.cap {
			delete(l.nodes, l.last.key)
			l.removeLast()
		} else {
			l.len++
		}
		node := &doublyLinkedNode{
			prev: nil,
			next: nil,
			key:  key,
			val:  value,
		}
		l.nodes[key] = node
		l.insertToFirst(node)
	}
}

func (l *LRUCache) removeLast() {
	if l.last.prev != nil {
		//双向链表长度>1
		l.last.prev.next = nil
	} else {
		//双向链表长度 == 1,first == last
		l.first = nil
	}
	l.last = l.last.prev
}
func (l *LRUCache) moveToFirst(node *doublyLinkedNode) {
	switch node {
	case l.first:
		return
	case l.last:
		l.removeLast()
	default:
		//在双向链中，删除 node 节点
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	// 策略是
	// 如果是移动node
	// 先删除,再插入
	l.insertToFirst(node)
}

func (l *LRUCache) insertToFirst(node *doublyLinkedNode) {
	if l.last == nil {
		//空的双向链表
		l.last = node
	} else {
		node.next = l.first
		l.first.prev = node
	}
	l.first = node
}
```

### GitHub
- [源码在这里](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0146.lru-cache/src/lru_cache.go)
- 项目中会提供各种数据结构及算法的Golang实现,以及 LeetCode 解法
