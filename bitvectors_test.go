package main

import "testing"

func Test_palindromePermutationBitVector(t *testing.T) {
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
			if got := palindromePermutationBitVector(tt.args.str); got != tt.want {
				t.Errorf("palindromePermutationBitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
