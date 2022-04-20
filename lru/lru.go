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
	var val = -1
	el, ok := this.hash[key]
	if ok {
		this.list.MoveToFront(el)
		val = el.Value.(Pair).Value
	}
	return val
}

func (this *LRUCache) Put(key int, value int) {
	el := &list.Element{Value: Pair{Key: key, Value: value}}
	this.hash[key] = el
	if this.capacity < len(this.hash) {
		back := this.list.Back()
		pair := back.Value.(*Pair)
		delete(this.hash, pair.Key)
		this.list.Remove(back)
	}
	this.list.PushFront(el)
}
