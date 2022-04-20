package main

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "ex 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: nil,
			// [-1]
			// [-1]
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVL_FindAt(t *testing.T) {
	type fields struct {
		add []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "ex",
			fields: fields{
				add: []int{3, 2, 1, 5},
			},
			args: args{i: 2},
			want: 3,
		},
		{
			name: "ex1",
			fields: fields{
				add: []int{3, 2, 1, 5},
			},
			args: args{i: 3},
			want: 5,
		},
		{
			name: "ex2",
			fields: fields{
				add: []int{1, 2},
			},
			args: args{i: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AVL{}
			for _, val := range tt.fields.add {
				a.Insert(val)
			}
			if got := a.FindAt(tt.args.i); got != tt.want {
				t.Errorf("AVL.FindAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAVL_Insert(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ex",
			args: args{
				vals: []int{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AVL{}
			for _, v := range tt.args.vals {
				a.Insert(v)
				a.Print()
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want *TreeNode
	}{
		{
			name: "",
			this: &Codec{},
			args: args{
				data: "1,2,3",
			},
			want: nil,
		},
		{
			name: "",
			this: &Codec{},
			args: args{
				data: "1,2,3,#,#,4,5,6,7",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Codec.deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_serialize(t *testing.T) {
	var c = &Codec{}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want string
	}{
		{
			name: "",
			this: c,
			args: args{
				root: c.deserialize("1,2,3"),
			},
			want: "1,2,3,#,#,#,#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorder(t *testing.T) {
	type args struct {
		n *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				n: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 8}},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 6, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 10}},
						Right: &TreeNode{Val: 7},
					}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 8 4 5 2 9 6 10 6 7 3 1
			// fmt.Print("postorder: ")
			// postorder(tt.args.n)
			// fmt.Println()
			// postorder2(tt.args.n)
			// inorderRecur(tt.args.n) // 8 4 2 5 1 9 6 10 3 7
			// inorder(tt.args.n) // 8 4 2 5 1 9 6 10 3 7
			// fmt.Print("inorder: ")
			// inorder(tt.args.n)
			// fmt.Println()
			preorderRecur(tt.args.n) // 1 2 4 8 5 3 6 9 10 7
			preorder(tt.args.n)
		})
	}
}

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"oath", "eat"},
		},
		{
			args: args{
				board: [][]byte{
					{'o', 'a', 'b', 'n'},
					{'o', 't', 'a', 'e'},
					{'a', 'h', 'k', 'r'},
					{'a', 'f', 'l', 'v'},
				},
				words: []string{"oa", "oaa"},
			},
			want: []string{"oa", "oaa"},
		},
		{
			args: args{
				board: [][]byte{
					{'a', 'a'},
				},
				words: []string{"aa"},
			},
			want: []string{"aa"},
		},
		{
			args: args{
				board: [][]byte{
					{'a', 'a'},
				},
				words: []string{"aaa"},
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
