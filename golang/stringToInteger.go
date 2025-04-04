package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i, result, sign := 0, 0, 1
	n := len(s)

	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}

func callMyAtoi() {
	fmt.Println(myAtoi("123"))
	fmt.Println(myAtoi("-123"))
	fmt.Println(myAtoi("0123"))
	fmt.Println(myAtoi("-0123"))
	fmt.Println(myAtoi("-0123bc1"))
	fmt.Println(myAtoi("0123bc1"))
}
