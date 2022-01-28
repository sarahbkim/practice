package main

import (
	"testing"
)

func Test_sockMerchant(t *testing.T) {
	type args struct {
		n  int32
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "basic",
			args: args{
				n:  9,
				ar: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sockMerchant(tt.args.n, tt.args.ar); got != tt.want {
				t.Errorf("sockMerchant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countingValleys(t *testing.T) {
	type args struct {
		steps int32
		path  string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				steps: 8,
				path:  "DDUUUUDD",
			},
			want: 1,
		},
		{
			args: args{
				steps: 8,
				path:  "UDDDUDUU",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingValleys(tt.args.steps, tt.args.path); got != tt.want {
				t.Errorf("countingValleys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jumpingOnClouds(t *testing.T) {
	type args struct {
		c []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				c: []int32{0, 0, 0, 0, 1, 0},
			},
			want: 3,
		},
		{
			args: args{
				c: []int32{0, 0, 0, 1, 0, 0},
			},
			want: 3,
		},
		{
			args: args{
				c: []int32{0, 0, 1, 0, 0, 1, 0},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpingOnClouds(tt.args.c); got != tt.want {
				t.Errorf("jumpingOnClouds() = %v, want %v", got, tt.want)
			}
		})
	}
}
