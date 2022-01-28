package main

import (
	"math"
	"sort"
	"strings"
)

// main concept: previous calculations can be kept as a single int.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// NOTE: i dont need this subtotal
	subtotals := make([]int, len(nums))

	// start with first
	subtotals[0] = nums[0]
	maxsum := nums[0]

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		sum := subtotals[i-1] + curr
		if curr > sum {
			subtotals[i] = curr
		} else {
			subtotals[i] = sum
		}
		if subtotals[i] > maxsum {
			maxsum = subtotals[i]
		}
	}
	return maxsum
}

func maxSubArrayLessMem(nums []int) int {
	var currSum float64
	var maxSum = math.Inf(-1)

	for i := 0; i < len(nums); i++ {
		currSum = math.Max(float64(nums[i]), currSum + float64(nums[i]))
		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return int(maxSum)
}

// main concept: prefix sum - calculate ranges quickly
func productExceptSelf(nums []int) []int {
	var prod = make([]int, len(nums))
	var l = 1
	for i := 0; i < len(nums); i++ {
		prod[i] = l
		l *= nums[i]
	}
	var r = 1
	for i := len(nums)-1; i >= 0; i-- {
		prod[i] = r * prod[i]
		r *= nums[i]
	}
	return prod
}

func maxProduct(nums []int) int {
	var product = nums[0]
	var min = nums[0]
	var max = nums[0]

	for i := 1; i < len(nums); i++ {
		min1 := min * nums[i]
		max1 := max * nums[i]

		min = int(math.Min(float64(max1), math.Min(float64(min1), float64(nums[i]))))
		max = int(math.Max(float64(max1), math.Max(float64(min1), float64(nums[i]))))
		product = int(math.Max(float64(product), float64(max)))
	}
	return product
}

func findMin(nums []int) int {
	if nums == nil {
		return -1
	}

	return bsearchMin(nums, 0, len(nums)-1)
}

func bsearchMin(nums []int, s, e int) int {
	if nums == nil {
		return -1
	}

	if s >= e {
		return nums[s]
	}

	if nums[s] < nums[e] {
		return nums[s]
	}
	m := (s + e) / 2
	l := bsearchMin(nums, s, m)
	r := bsearchMin(nums, m+1, e)
	if l < r {
		return l
	}
	return r
}

func search(nums []int, target int) int {
	return rotatedSearchHelper(nums, 0, len(nums)-1, target)
}

func rotatedSearchHelper(nums []int, s, e, t int) int {
	if e >= s {
		m := (s + e) / 2
		if nums[m] == t{ return m }
		// sorted side
		if nums[s] < nums[e] {

			if nums[m] < t {
				return rotatedSearchHelper(nums, m+1, e, t)
			} else {
				return rotatedSearchHelper(nums, s, m, t)
			}
		} else if s < e { // be careful here! make sure no infinite recursion
			// NOTE: shouldn't need to search both sides
			l := rotatedSearchHelper(nums, s, m, t)
			r := rotatedSearchHelper(nums, m+1, e, t)
			if l > -1 {
				return l
			}
			if r > -1 {
				return r
			}
		}
	}
	return -1
}

// runs o(n^2) time
func threeSum(nums []int) [][]int {
	var ans [][]int
	// sort!
	sort.Ints(nums)

	// optimize
	if len(nums) == 0 || nums[0] > 0 || nums[len(nums)-1] < 0 {
	return nil
	}

	for i := 0; i < len(nums) - 2; i++ {
		// remove dupes
		if i == 0 || nums[i] != nums[i-1] {
			var lo = i + 1
			var hi = len(nums)-1
			sum := 0 - nums[i]
			// bidirectional sweep of array - 2sum
			for lo < hi {
				if nums[lo] + nums[hi] + nums[i] == 0 {
					ans = append(ans, []int{nums[i], nums[lo], nums[hi]})
					// remove dupes
					for lo < hi && nums[lo] == nums[lo+1] { lo++ }
					// remove dupes
					for lo < hi && nums[hi] == nums[hi-1] { hi-- }
					lo++
					hi--
				} else if nums[lo] + nums[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}

	return ans
}

// 1.3 - URL-ify replaces all of the spaces with
// '%20', assumes that str []byte contains enough space
// trueLen is the length of the actual string str, without
// the extra buffer appended
func urlify(str []byte, trueLen int) string {
	var spaces int
	for _, c := range str {
		if c == ' ' {
			spaces++
		}
	}

	bs := make([]byte, spaces * 2 + trueLen)
	copy(bs, str)

	var index = (spaces * 2) + trueLen
	for i := trueLen-1; i >= 0; i--{
		if bs[i] == ' ' {
			bs[index-1] = '0'
			bs[index-2] = '2'
			bs[index-3] = '%'
			index -= 3
		} else {
			bs[index-1] = bs[i]
			index--
		}
	}
	return string(bs)
}

// 1.4 - Check if it is a permutation of a palindrome
func palindromePermutation(str string) bool {
	if len(str) == 1 { return true  }
	var dict = map[string]int{}
	for _, c := range str {
		// todo: make lowercase
		if c == ' ' { continue }
		char := strings.ToLower(string(c))
		dict[char]++
	}
	var odds bool
	for _, counts := range dict {
		if counts % 2 != 0 {
			if odds { return false }
			odds = true
		}
	}
	return true
}

// 1.5
func oneEditAway(str1, str2 string) bool {
	if len(str1) == len(str2) {
		return checkOneReplace(str1, str2)
	} else if len(str1) - 1 == len(str2)  {
		return checkOneInsert(str2, str1)
	} else if len(str2) - 1 == len(str1) {
		return checkOneInsert(str1, str2)
	}
	return false
}

func checkOneReplace(str1, str2 string) bool {
	var foundDiff bool
	for i, j := 0, 0; i < len(str1) && j < len(str2); i, j = i + 1, j + 1 {
		if str1[i] != str2[j] {
			if foundDiff { return false }
			foundDiff = true
		}
	}
	return true
}

func checkOneInsert(str1, str2 string) bool {
	var insertCount int
	for i, j := 0, 0; i < len(str1) && j < len(str2); {
		if str1[i] != str2[j] {
			if insertCount > 1 { return false }
			insertCount++
			// increment the pointer for longer string
			if len(str1) < len(str2) {
				j++
			} else {
				i++
			}
			continue
		}
		i++
		j++
	}

	return true
}

func rotateMatrix(matrix [][]int) [][]int {
	// reverse
	for i := 0; i < len(matrix); i++ {
		for j, k := 0, len(matrix[i])-1; j < k; j, k = j + 1, k - 1 {
			matrix[j][i], matrix[k][i] = matrix[k][i], matrix[j][i]
		}
	}
	// diagonally swap
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func zeroMatrix(A [][]int) [][]int {
	var rows []int
	var cols []int

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, i)
			}
		}
	}

	for _, row := range rows {
		A[row] = make([]int, len(A[row]))
	}
	for _, col := range cols {
		for i := 0 ; i < len(A); i++ {
			A[i][col] = 0
		}
	}
	return A
}