package main

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
