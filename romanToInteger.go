package main

func romanToInt(s string) int {
	num := 0
	myMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	num += myMap[string(s[0])]
	for i := 1; i < len(s); i++ {
		if val, ok := myMap[string(s[i])]; ok {
			if myMap[string(s[i])] > myMap[string(s[i-1])] {
				num += val - 2*myMap[string(s[i-1])]
			} else {
				num += val
			}
		}
	}
	return num
}
