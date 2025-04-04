package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	foundCharLastAt := make(map[byte]int)
	maxLength := 0

	for i, j := 0, 0; j < len(s); j++ {
		if lastIndex, exists := foundCharLastAt[s[j]]; exists {
			i = max(i, lastIndex+1)
		}
		maxLength = max(maxLength, j-i+1)
		foundCharLastAt[s[j]] = j
	}

	return maxLength
}

func callLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))   //3
	fmt.Println(lengthOfLongestSubstring("abcd"))   //4
	fmt.Println(lengthOfLongestSubstring("abcdea")) //5
	fmt.Println(lengthOfLongestSubstring("pwwkew")) //3
}
