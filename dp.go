package main

import (
	"math"
	"sort"
)

func combinationSumClosure(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	tmp := []int{}
	var dfs func(int, int)
	dfs = func(currTotal int, idx int) {
		if currTotal == target {
			copied := make([]int, len(tmp))
			copy(copied, tmp)
			res = append(res, copied)
			return
		}
		if idx == len(candidates) {
			return
		}
		for i := idx; i < len(candidates); i++ {
			if currTotal+candidates[i] > target {
				break
			}
			tmp = append(tmp, candidates[i])
			dfs(currTotal+candidates[i], i)
			tmp = tmp[0 : len(tmp)-1]
		}
	}
	dfs(0, 0)
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	return comboSum(candidates, 0, 0, []int{}, target)
}

func comboSum(candidates []int, i int, currSum int, list []int, target int) [][]int {
	if currSum == target {
		return [][]int{list}
	}
	if currSum > target {
		return nil
	}

	var ans = [][]int{}
	// i, j will ensure that each tree gets same subset of candidates i to j
	for ; i < len(candidates); i++ {
		sum := currSum + candidates[i]

		// make copies of this list, or things get re-used
		var buf = make([]int, len(list))
		copy(buf, list)
		buf = append(buf, candidates[i])

		res := comboSum(candidates, i, sum, buf, target)
		ans = append(ans, res...)
	}
	return ans
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var memo = make([]int, len(nums)+1)
	memo[1] = nums[0]

	// i - 2 and i - 1
	for i := 1; i < len(nums); i++ {
		var val = nums[i]
		memo[i+1] = int(math.Max(float64(memo[i]), float64(memo[i-1]+val)))
	}
	return memo[len(nums)]
}

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	var a = 1
	var b = 2
	if n == 1 {
		return a
	}
	if n == 2 {
		return b
	}

	var ways int
	for i := 2; i < n; i++ {
		ways = a + b
		a = b
		b = ways
	}
	return ways
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var x = make([]float64, amount+1)
	for i := 0; i < len(x); i++ {
		x[i] = math.Inf(1)
	}

	for i := 1; i < len(x); i++ {
		for _, c := range coins {
			if c > i {
				x[i] = math.Min(x[i], math.Inf(1))
			} else if c == i {
				x[i] = math.Min(x[i], 1)
			} else {
				rem := i - c
				x[i] = math.Min(1+x[rem], x[i])
			}
		}
	}

	if x[amount] == math.Inf(1) {
		return -1
	}
	return int(x[amount])
}

func numDecodings(s string) int {
	var dp = make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	for i := 2; i < len(dp); i++ {
		if s[i-1] >= '1' && s[i-1] <= '9' {
			dp[i] += dp[i-1]
		}
		if s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func wordBreak(s string, wordDict []string) bool {
	var dp = make([]bool, len(s)+1)
	dp[0] = true

	var dict = make(map[string]bool, len(wordDict))

	for _, word := range wordDict {
		dict[word] = false
	}

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			substr := s[j:i]
			if _, ok := dict[substr]; ok && dp[j] {
				dict[substr] = true
				dp[i] = true
			}
		}
	}
	return dp[len(dp)-1]
}

func longestOnes(nums []int, k int) int {
	var longest, zeroes int
	var i, j int
	for ; j < len(nums); j++ {
		if nums[j] == 0 {
			zeroes++
		}
		for zeroes > k {
			if nums[i] == 0 {
				zeroes--
			}
			i++
		}
		curr := j - i + 1
		longest = max(longest, curr)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	var longest int
	var i, j int
	var answers = make(map[byte]int, 2)

	var maxOccur int
	for ; j < len(answerKey); j++ {
		answers[answerKey[j]]++
		maxOccur = max(maxOccur, answers[answerKey[j]])
		for j-i+1-maxOccur > k {
			answers[answerKey[i]]--
			i++
		}
		longest = max(longest, j-i+1)
	}
	return longest
}

// func longestPalindrome(s string) string {
// 	var longest string
// 	var cache = map[[2]int]bool{}
// 	for i := 0; i < len(s); i++ {
// 		for j := 0; j < len(s); j++ {
// 			if i == j || (s[i] == s[j] && cache[[2]int{i + 1, j - 1}]) {
// 				cache[[2]int{i, j}] = true
// 				if len(longest) < j-i+1 {
// 					longest = s[i : j+1]
// 				}
// 			} else {
// 				cache[[2]int{i, j}] = false
// 			}
// 		}
// 	}
// 	return longest
// }

func longestPalindrome(s string) string {
	var longest string
	var cache = map[[2]int]bool{}
	var size int
	for size < len(s) {
		for i := 0; i < len(s)-size; i++ {
			var lo = i
			var hi = i + size
			valid, _ := cache[[2]int{lo + 1, hi - 1}]
			if lo == hi || (s[lo] == s[hi] && valid) {
				if len(longest) < len(s[lo:hi+1]) {
					longest = s[lo : hi+1]
				}
				cache[[2]int{lo, hi}] = true
			} else {
				cache[[2]int{lo, hi}] = false
			}
		}
		size++
	}
	return longest
}
