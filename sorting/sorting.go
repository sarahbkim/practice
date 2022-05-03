package sorting

import (
	"math"
)

func directAccessSort(nums []int) []int {
	// get max val - O(n)
	var max = math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	// O(n)
	var directAccessArr = make([]*int, max+1)
	for i := 0; i < len(nums); i++ {
		(directAccessArr[nums[i]]) = &nums[i]
	}
	// O(n)
	var sorted []int
	for i := 0; i < len(directAccessArr); i++ {
		if directAccessArr[i] != nil {
			sorted = append(sorted, *directAccessArr[i])
		}
	}
	return sorted
}

func quicksort(A []int, lo, hi int) {
	if lo < hi {
		s, e := partition(A, lo, hi)
		quicksort(A, lo, s)
		quicksort(A, e, hi)
	}
}

func partition(A []int, lo, hi int) (int, int) {
	var i, j, k = lo, hi, lo
	var mid = (hi + lo) / 2
	var pivot = A[mid]
	for k < j {
		if A[k] < pivot {
			A[i], A[k] = A[k], A[i]
			k++
			i++
		} else if A[k] == pivot {
			k++
		} else {
			j--
			A[k], A[j] = A[j], A[k]
		}
	}
	return i, k
}

func findKthSmallest(A []int, k int) int {
	if len(A) == 1 {
		return A[0]
	}
	p, _ := partition(A, 0, len(A))
	if p == k {
		return A[p]
	}
	left := A[:p+1]
	right := A[p+2:]
	if k < len(left) {
		return findKthSmallest(left, k-len(left))
	}
	return findKthSmallest(right, k-len(right))
}
