package sorting

import "math"

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
