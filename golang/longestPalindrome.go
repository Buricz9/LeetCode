package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	maxStr := s[0:1]

	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > len(maxStr) {
				maxStr = s[left : right+1]
			}
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > len(maxStr) {
				maxStr = s[left : right+1]
			}
			left--
			right++
		}
	}
	return maxStr
}

func callLongestPalindromeSubstring() {

	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("aa"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("ac"))
}
