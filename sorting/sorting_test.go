package sorting

import (
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic test i",
			args: args{
				A: []int{3, 2, 4, 1, 2, 2, 3},
			},
			want: []int{1, 2, 2, 2, 3, 3, 4},
		},
		{
			name: "basic test ii",
			args: args{
				A: []int{100, 21, 40, 97, 53, 9, 25, 105, 99, 8, 45, 10},
			},
			want: []int{8, 9, 10, 21, 25, 40, 45, 53, 97, 99, 100, 105},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args.A, 0, len(tt.args.A))
			if !reflect.DeepEqual(tt.args.A, tt.want) {
				t.Errorf("failed %s. got %v, want %v", tt.name, tt.args.A, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		expectedA []int
	}{
		{
			name: "basic test",
			args: args{
				A: []int{3, 2, 5, 1, 4},
			},
			want:      2,
			expectedA: []int{3, 2, 1, 4, 5},
		},
		{
			name: "basic test - ii",
			args: args{
				A: []int{3, 5, 2, 1, 4},
			},
			want:      2,
			expectedA: []int{1, 2, 5, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partition(tt.args.A, 0, len(tt.args.A))
			if tt.expectedA != nil {
				if !reflect.DeepEqual(tt.args.A, tt.expectedA) {
					t.Errorf("partition() = %v, want %v", tt.args.A, tt.expectedA)
				}
			}
		})
	}
}

func Test_directAccessSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := directAccessSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("directAccessSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthSmallest(t *testing.T) {
	type args struct {
		A []int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "basic",
			args: args{
				A: []int{100, 21, 40, 97, 53, 9, 25, 105, 99, 8, 45, 10},
				k: 7,
			},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthSmallest(tt.args.A, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
