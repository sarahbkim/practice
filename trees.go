package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// generate a map for the index of a node
	var idx = make(map[int]int, len(inorder))
	for i, val := range inorder {
		idx[val] = i
	}
	var preorderIdx int
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		rootVal := preorder[preorderIdx]

		n := &TreeNode{Val: rootVal}
		preorderIdx++

		n.Left = build(l, idx[rootVal]-1)
		n.Right = build(idx[rootVal]+1, r)
		return n
	}
	node := build(0, len(preorder)-1)
	return node
}

func buildTreeLinear(preorder []int, inorder []int) *TreeNode {
	var build func(stopVal *int) *TreeNode
	build = func(stopVal *int) *TreeNode {
		if len(inorder) > 0 && (stopVal == nil || inorder[0] != *(stopVal)) {
			rootVal := preorder[0]
			preorder = preorder[1:]
			root := &TreeNode{Val: rootVal}
			root.Left = build(&root.Val)
			inorder = inorder[1:]
			root.Right = build(stopVal)
			return root
		}
		return nil
	}

	return build(nil)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	var freq int
	var prev *int
	var s = []*TreeNode{}

	var maxFreq int
	var modes = []int{}

	for root != nil || len(s) > 0 {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}

		// pop from stack
		root = s[len(s)-1]
		s = s[:len(s)-1]

		if prev == nil || (root.Val == *(prev)) {
			freq++
			if freq > maxFreq {
				maxFreq = freq
				modes = []int{root.Val}
			} else if freq == maxFreq {
				modes = append(modes, root.Val)
			}
		}

		prev = &root.Val
		root = root.Right
	}
	return modes
}

type Iter interface {
	FindAt(i int) int
	Insert(val int)
}

type n struct {
	counts int
	val    int
	left   *n
	right  *n
}

type AVL struct {
	root *n
}

func (a *AVL) FindAt(i int) int {
	var node = a.root
	if node != nil && node.counts < i {
		return -1
	}
	var found int
	for node != nil {
		var leftsize int
		if node.left != nil {
			leftsize = node.left.counts
		}
		if i < leftsize {
			node = node.left
		} else if i > leftsize {
			i = i - leftsize - 1
			node = node.right
		} else {
			found = node.val
			break
		}
	}
	return found
}

func (a *AVL) Insert(val int) {
	var prev *n
	var node = a.root
	if node == nil {
		node = &n{val: val, counts: 1}
		a.root = node
		return
	}
	for node != nil {
		node.counts++
		prev = node
		if val < node.val {
			node = node.left
		} else {
			node = node.right
		}
	}
	if val < prev.val {
		prev.left = &n{val: val, counts: 1}
	} else {
		prev.right = &n{val: val, counts: 1}
	}
	// TODO: balance
}

func (a *AVL) Print() {
	if a.root == nil {
		return
	}
	var levels = [][]*n{[]*n{a.root}}
	for len(levels) > 0 {
		level := levels[0]
		levels = levels[1:]

		var newLevel = []*n{}
		for i := 0; i < len(level); i++ {
			node := level[i]
			if node == nil {
				fmt.Print(nil)
			} else {
				fmt.Printf("curr node val:{%d}, counts:{%d}", node.val, node.counts)
				newLevel = append(newLevel, node.left)
				newLevel = append(newLevel, node.right)
			}
		}
		if len(newLevel) > 0 {
			levels = append(levels, newLevel)
		}
		fmt.Println()
	}
}

func NewAVL() Iter {
	return &AVL{root: nil}
}

// use a balanced binary tree
type MedianFinder struct {
	iter   Iter
	ncount int
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{
		iter:   NewAVL(),
		ncount: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.ncount++
	this.iter.Insert(num)

}

func (this *MedianFinder) FindMedian() float64 {
	mid := this.ncount / 2
	if this.ncount%2 == 0 {
		mids := this.iter.FindAt(mid) + this.iter.FindAt(mid-1)
		return float64(mids) / 2
	}
	return float64(this.iter.FindAt(mid))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var vals []string
	var myFn func(node *TreeNode)
	myFn = func(node *TreeNode) {
		if node == nil {
			vals = append(vals, "#")
		} else {
			vals = append(vals, strconv.Itoa(node.Val))
			myFn(node.Left)
			myFn(node.Right)
		}
	}
	myFn(root)
	return strings.Join(vals, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var myFn func() *TreeNode
	vals := strings.Split(data, ",")
	myFn = func() *TreeNode {
		if len(vals) == 0 {
			return nil
		}
		pop := vals[0]
		vals = vals[1:]

		if pop == "#" {
			return nil
		}
		v, _ := strconv.Atoi(pop)
		node := &TreeNode{Val: v}
		node.Left = myFn()
		node.Right = myFn()
		return node
	}
	return myFn()
}

func inorder(n *TreeNode) {
	var q = []*TreeNode{}
	for len(q) > 0 || n != nil {
		for n != nil {
			q = append(q, n)
			n = n.Left
		}
		n = q[len(q)-1]
		q = q[0 : len(q)-1]
		fmt.Print(n.Val, ",")

		n = n.Right
	}
	fmt.Println()
}

func inorderRecur(n *TreeNode) {
	if n != nil {
		inorderRecur(n.Left)
		fmt.Print(n.Val, ",")
		inorderRecur(n.Right)
	}
}

func preorder(n *TreeNode) {
	var q = []*TreeNode{}
	var s = []*TreeNode{}
	q = append(q, n)
	for len(q) > 0 || len(s) > 0 {
		for len(q) > 0 {
			n = q[0]
			q = q[1:]
			fmt.Print(n.Val, ",")
			s = append(s, n)
			if n.Left != nil {
				q = append(q, n.Left)
			}
		}
		n = s[len(s)-1]
		s = s[0 : len(s)-1]
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	fmt.Println()
}

func preorderRecur(n *TreeNode) {
	if n != nil {
		fmt.Print(n.Val, ",")
		preorderRecur(n.Left)
		preorderRecur(n.Right)
	}
}

func postorder(n *TreeNode) {
	var q = []*TreeNode{}
	var pre *TreeNode
	for len(q) > 0 || n != nil {
		for n != nil {
			q = append(q, n)
			n = n.Left
		}
		n = q[len(q)-1]
		if n.Right == nil || n.Right == pre {
			fmt.Print(n.Val, ",")
			q = q[0 : len(q)-1]
			pre = n
			n = nil
		} else {
			n = n.Right
		}
	}
	fmt.Println()
}

func postorder2(n *TreeNode) {
	var s = []*TreeNode{}
	var r = []*TreeNode{}
	for n != nil || len(s) > 0 {
		for n != nil {
			if n.Right != nil {
				r = append(r, n.Right)
			}
			s = append(s, n)
			n = n.Left
		}
		n = s[len(s)-1]
		if len(r) > 0 && n.Right == r[len(r)-1] {
			n = r[len(r)-1]
			r = r[0 : len(r)-1]
		} else {
			fmt.Print(n.Val, ",")
			s = s[0 : len(s)-1]
			n = nil
		}
	}
	fmt.Println()
}

func postOrderRecur(n *TreeNode) {
	if n == nil {
		return
	}
	postOrderRecur(n.Left)
	postOrderRecur(n.Right)
	fmt.Print(n.Val, ",")
}

func findWords(board [][]byte, words []string) []string {
	t := NewTrie()
	for _, w := range words {
		t.Insert(w)
	}

	var prefix bytes.Buffer
	var found = map[string]struct{}{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			for _, k := range dfs(prefix, i, j, board, &t) {
				found[k] = struct{}{}
			}
		}
	}

	var ans = make([]string, len(found))
	var i int
	for k := range found {
		ans[i] = k
		i++
	}

	return ans
}

func dfs(prefix bytes.Buffer, i, j int, board [][]byte, t *Trie) []string {
	var ans []string
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return ans
	}
	tmp := board[i][j]
	prefix.WriteByte(tmp)
	board[i][j] = '#'
	if t.Search(prefix.String()) {
		ans = append(ans, prefix.String())
	}
	if t.StartsWith(prefix.String()) {
		ans = append(ans, dfs(prefix, i+1, j, board, t)...)
		ans = append(ans, dfs(prefix, i-1, j, board, t)...)
		ans = append(ans, dfs(prefix, i, j+1, board, t)...)
		ans = append(ans, dfs(prefix, i, j-1, board, t)...)
	}
	board[i][j] = tmp
	return ans
}
