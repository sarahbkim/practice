package lru

import (
	"container/list"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	type fields struct {
		hash     map[int]*list.Element
		list     *list.List
		capacity int
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &LRUCache{
				hash:     tt.fields.hash,
				list:     tt.fields.list,
				capacity: tt.fields.capacity,
			}
			this.Put(tt.args.key, tt.args.value)
		})
	}
}
