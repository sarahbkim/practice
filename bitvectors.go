package main

import (
	"unicode"
)

// 1.4 - Check if it is a permutation of a palindrome
func palindromePermutationBitVector(str string) bool {
	bitVector := createBitVector(str)
	return bitVector == 0 || checkOneBitSet(bitVector)
}

func createBitVector(str string) int {
	var vector int
	for _, c := range str {
		i := unicode.ToLower(c) - 'a'
		vector = toggle(vector, i)
	}
	return vector
}

func toggle(bitVector int, index rune) int {
	if index < 0 { return bitVector }
	mask := 1 << index
	if (bitVector & mask) == 0 {
		bitVector |= mask
	} else {
		bitVector &= ^mask
	}
	return bitVector
}

func checkOneBitSet(bitVector int) bool {
	sub := bitVector - 1
	return bitVector & sub == 0
}
