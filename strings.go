package main

import (
	"unicode"
)

// assuming english alphabet 26 chars
func isAnagram(s string, t string) bool {
	var alpha = make([]int, 26)
	for _, r := range s {
		alpha[r-'a']++
	}
	for _, r := range t {
		if alpha[r-'a'] == 0 {
			return false
		}
		alpha[r-'a']--
	}
	for _, count := range alpha {
		if count > 0 {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	var l = 0
	var r = len(s) - 1
	for l < r {
		for l < len(s) && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r >= 0 && !isAlphaNum(rune(s[r])) {
			r--
		}
		if l < len(s) && r >= 0 {
			left := unicode.ToLower(rune(s[l]))
			right := unicode.ToLower(rune(s[r]))
			if left != right {
				return false
			}
			l++
			r--
		}
	}
	return true
}

func isAlphaNum(char rune) bool {
	return unicode.IsDigit(char) || unicode.IsLetter(char)
}
