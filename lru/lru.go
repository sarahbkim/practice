package lru

import (
	"container/list"
)

type Pair struct {
	Key   int
	Value int
}

type LRUCache struct {
	hash     map[int]*list.Element
	list     *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hash:     make(map[int]*list.Element, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hash[key]; ok {
		this.list.MoveToFront(node)
		return node.Value.(Pair).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	el, ok := this.hash[key]
	if ok {
		this.list.MoveToFront(el)
		pair := el.Value.(Pair)
		pair.Value = value
	} else {
		if this.capacity == this.list.Len() {
			last := this.list.Back()
			pair := last.Value.(Pair)
			delete(this.hash, pair.Key)
			this.list.Remove(last)
		} else {
			node := &list.Element{Value: Pair{Key: key, Value: value}}
			this.hash[key] = node
			this.list.PushFront(el)
		}

	}
}
