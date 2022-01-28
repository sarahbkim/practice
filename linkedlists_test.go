package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_removeDupes(t *testing.T) {
	type args struct {
		n *node
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "basic test",
			args: args{
				n: &node{val: 1, next: &node{val: 2, next: &node{val: 1}}},
			},
			expected: "12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeDupes(tt.args.n)
			n := tt.args.n
			if !reflect.DeepEqual(n.toString(), tt.expected) {
				t.Fatalf("expected %s, got %s", tt.expected, n.toString())
			}
		})
	}
}

func Test_kthLastElement(t *testing.T) {
	type args struct {
		k int
		n *node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				k: 2,
				n: &node{val: 1, next: &node{val: 2, next: &node{val: 3, next: &node{val: 4}}}},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				k: 4,
				n: &node{val: 1, next: &node{val: 2, next: &node{val: 3, next: &node{val: 4}}}},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				k: 3,
				n: &node{val: 1, next: &node{val: 2, next: &node{val: 3, next: &node{val: 4}}}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLastElement(tt.args.k, tt.args.n); !reflect.DeepEqual(got.val, tt.want) {
				t.Errorf("kthLastElement() = %v, want %v", got.val, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		n   *node
		val int
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "1",
			args: args{
				n:   &node{val: 2, next: &node{val: 5, next: &node{val: 1, next: &node{val: 0}}}},
				val: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.n, tt.args.val)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
			fmt.Println(got.toString())
		})
	}
}

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: generateList([]int{1, 2, 3, 4}),
			},
			want: generateList([]int{1, 4, 2, 3}),
		},
		{
			name: "2",
			args: args{
				head: generateList([]int{1, 2, 3, 4, 5}),
			},
			want: generateList([]int{1, 5, 2, 4, 3}),
		},
		{
			name: "3",
			args: args{
				head: generateList([]int{1, 2, 3}),
			},
			want: generateList([]int{1, 3, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got.toString(), tt.want.toString())
			}
		})
	}
}
