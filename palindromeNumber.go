package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	number := 0

	for x > 0 {
		number = number*10 + x%10
		x = x / 10
	}
	return original == number
}
