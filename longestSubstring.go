package main

func lengthOfLongestSubstring(s string) int {

	mapString := make(map[string]int)
	count := 0
	longest := 0
	for i := 0; i < len(s); i++ {
		if _, ok := mapString[string(s[i])]; !ok {
			count++
			mapString[string(s[i])] = i
			if count > longest {
				longest = count
			}
		} else {
			if i != len(s)-1 {
				count = 0
				i = mapString[string(s[i])]
				mapString = make(map[string]int)
			}
		}
	}
	return longest
}
