package main

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type node struct {
	next *node
	val  int
}

func (n *node) toString() string {
	var curr = n
	var b bytes.Buffer
	for curr != nil {
		b.WriteString(strconv.Itoa(curr.val))
		curr = curr.next
	}
	return b.String()
}

func removeDupes(n *node) {
	curr := n
	for curr != nil {
		r := curr
		for r.next != nil {
			if r.next.val == curr.val {
				r.next = r.next.next
			} else {
				r = r.next
			}
		}
		curr = curr.next
	}
}

func kthLastElement(k int, n *node) *node {
	var fast = n
	var slow = n
	for i := 0; i < k; i++ {
		fast = fast.next
	}
	for fast != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow
}

func partition(n *node, val int) *node {
	var tail = &node{}
	var head = &node{}
	var a = head
	var b = tail
	var curr = n
	for curr != nil {
		if curr.val < val {
			head.next = &node{val: curr.val}
			head = head.next
		} else {
			tail.next = &node{val: curr.val}
			tail = tail.next
		}
		curr = curr.next
	}
	head.next = b.next
	return a.next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	var chain = reorderList(head.Next)
	var curr = head

	for curr != nil && chain != nil && chain.Next != nil {
		var tmp = chain.Next.Next
		curr.Next = chain.Next
		chain.Next.Next = chain
		chain.Next = tmp
		curr = curr.Next
	}

	return head
}

func generateList(vals []int) *ListNode {
	var head = &ListNode{}
	var curr = head
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head.Next
}

func (n *ListNode) toString() string {
	var curr = n
	var b bytes.Buffer
	for curr != nil {
		b.WriteString(strconv.Itoa(curr.Val))
		b.WriteString("->")
		curr = curr.Next
	}
	return b.String()
}
