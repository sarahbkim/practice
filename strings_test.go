package main

import (
	"testing"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "b",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "c",
			args: args{s: " "},
			want: true,
		},
		{
			name: "d",
			args: args{s: "a"},
			want: true,
		},
		{
			name: "e",
			args: args{s: ",."},
			want: true,
		},
		{
			name: "f",
			args: args{s: "0P"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("Test %s = isPalindrome() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
