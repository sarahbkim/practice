package main

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validTree(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTree(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("validTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
