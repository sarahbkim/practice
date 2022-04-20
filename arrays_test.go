package main

import (
	"reflect"
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2}, {-1, 0, 1},
			},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{},
			},
			want: nil,
		},
		{
			name: "example 3",
			args: args{
				nums: []int{0},
			},
			want: nil,
		},
		{
			name: "example 4",
			args: args{
				nums: []int{0, 0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "example 5",
			args: args{
				nums: []int{1, -1, -1, 0},
			},
			want: [][]int{{-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_urlify(t *testing.T) {
	type args struct {
		str     []byte
		trueLen int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic test",
			args: args{
				str:     []byte("Mr John Smith"),
				trueLen: 13,
			},
			want: "Mr%20John%20Smith",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlify(tt.args.str, tt.args.trueLen); got != tt.want {
				t.Errorf("urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_palindromePermutation(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "taco cat",
			args: args{
				str: "Tact Coa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePermutation(tt.args.str); got != tt.want {
				t.Errorf("palindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkOneInsert(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pale and ple",
			args: args{
				str1: "pale",
				str2: "ple",
			},
			want: true,
		},
		{
			name: "pales and ple",
			args: args{
				str1: "pales",
				str2: "pale",
			},
			want: true,
		},
		{
			name: "pales and aple",
			args: args{
				str1: "pales",
				str2: "aple",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOneInsert(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("checkOneInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oneEditAway(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				str1: "pale",
				str2: "ple",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				str1: "pales",
				str2: "pale",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				str1: "pale",
				str2: "bale",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				str1: "pale",
				str2: "bake",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "basic test",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zeroMatrix(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "basic",
			args: args{
				A: [][]int{
					{0, 1, 2},
					{3, 4, 5},
					{7, 8, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 4, 0},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroMatrix(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSmaller(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "basic",
		// 	args: args{
		// 		nums: []int{5, 4, 3, 2, 1},
		// 	},
		// 	want: []int{4, 3, 2, 1, 0},
		// },
		{
			name: "basic",
			args: args{
				nums: []int{26, 78, 27, 100, 33},
			},
			want: []int{0, 2, 0, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
