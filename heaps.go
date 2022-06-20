package main

import (
	"container/heap"
	"sort"
)

// TODO: can i use binary search? the inputs are sorted...
func kthSmallest(matrix [][]int, k int) int {
	var h = &MinHeap{}
	heap.Init(h)
	// put everything in a min heap
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			heap.Push(h, matrix[i][j])
		}
	}

	// pop kth smallest
	for k > 0 {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

// An IntHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	var pq rooms
	heap.Init(&pq)
	for i := 0; i < len(intervals); i++ {
		mtg := intervals[i]
		if pq.Len() == 0 {
			heap.Push(&pq, &room{value: mtg, meetingEnd: mtg[1]})
		} else {
			earliestMtgRoom := peek(&pq).(*room)
			if earliestMtgRoom.meetingEnd < mtg[0] {
				heap.Pop(&pq)
				heap.Push(&pq, &room{value: mtg, meetingEnd: mtg[1]})
			}

		}
	}
	return pq.Len()
}

type room struct {
	value      []int
	meetingEnd int
}

type rooms []*room

func (r rooms) Len() int { return len(r) }
func (r rooms) Less(i, j int) bool {
	return r[i].meetingEnd < r[j].meetingEnd
}
func (r rooms) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r *rooms) Push(x interface{}) {
	item := x.(*room)
	*r = append(*r, item)
}
func peek(r *rooms) interface{} {
	old := *r
	n := len(old)
	item := old[n-1]

	return item
}
func (r *rooms) Pop() interface{} {
	old := *r
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*r = old[0 : n-1]
	return item
}
