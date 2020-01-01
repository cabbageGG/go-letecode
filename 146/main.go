package main

import "fmt"

// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
// 进阶:

// 你是否可以在 O(1) 时间复杂度内完成这两种操作？

// 示例:

// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // 返回  1
// cache.put(3, 3);    // 该操作会使得密钥 2 作废
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得密钥 1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4

// 解法： 双链表 + hash
// 注意hash保存的值是链表节点。链表节点保存了key和value
// 注意构造虚拟头尾节点，方便操作

type ListNode struct {
	Key   int
	Value int
	Pre   *ListNode
	Next  *ListNode
}

type LRUCache struct {
	Capacity int
	Head     *ListNode
	Tail     *ListNode
	Map      map[int]*ListNode
}

func Constructor(capacity int) LRUCache {
	h := &ListNode{-1, -1, nil, nil}
	t := &ListNode{-1, -1, nil, nil}
	h.Next = t
	t.Pre = h
	cache := LRUCache{
		Capacity: capacity,
		Head:     h,
		Tail:     t,
		Map:      make(map[int]*ListNode, capacity),
	}
	return cache
}

func (this *LRUCache) insert(node *ListNode) {
	h := this.Head.Next // true head
	this.Head.Next = node
	node.Pre = this.Head
	node.Next = h
	h.Pre = node
}

func (this *LRUCache) remove(node *ListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Map[key]; ok {
		this.remove(v)
		this.insert(v)
		return v.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.Map[key]; ok {
		this.remove(v)
		this.insert(v)
		v.Value = value
	} else {
		if len(this.Map) >= this.Capacity {
			tail := this.Tail.Pre
			this.remove(tail)
			delete(this.Map, tail.Key)
		}
		node := &ListNode{key, value, nil, nil}
		this.insert(node)
		this.Map[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
