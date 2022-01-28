package main

import (
	"math"
)

func sockMerchant(n int32, ar []int32) int32 {
	var pairs int32
	hash := make(map[int32]bool, n)
	for i := 0; i < len(ar); i++ {
		if exists, ok := hash[ar[i]]; ok && exists {
			pairs++
			hash[ar[i]] = false
		} else {
			hash[ar[i]] = true
		}
	}
	return pairs
}

func countingValleys(steps int32, path string) int32 {
	var counts int32
	var currLevel int32
	for _, str := range path {
		if str == 'U' {
			currLevel++
		} else {
			currLevel--
		}
		// if we are back to 0 after seeing U
		if str == 'U' && currLevel == 0 {
			counts++
		}
	}
	return counts
}

func jumpingOnClouds(c []int32) int32 {
	steps := make([]float64, len(c))
	if c[0] == 1 {
		steps[0] = math.Inf(1)
	} else {
		steps[0] = 0
	}

	if c[1] == 1 {
		steps[1] = math.Inf(1)
	} else {
		steps[1] = 1
	}
	for i := 2; i < len(c); i++ {
		if c[i] == 1 {
			steps[i] = math.Inf(1)
		} else {
			steps[i] = math.Min(steps[i-1], steps[i-2]) + 1
		}
	}
	return int32(steps[len(c)-1])
}
