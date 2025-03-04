package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	const minInt, maxInt = -1 << 31, (1 << 31) - 1

	val := true
	if x < 0 {
		val = false
	}
	x = int(math.Abs(float64(x)))
	reversInt := 0

	for i := int(math.Log10(float64(x))); x != 0; i-- {
		num := x % 10 * int(math.Pow10(i))
		reversInt += num
		x /= 10

		if reversInt > maxInt || reversInt < minInt {
			return 0
		}
	}

	if !val {
		reversInt *= -1
	}

	if reversInt > maxInt || reversInt < minInt {
		return 0
	}

	return reversInt
}

func callReverse() {
	fmt.Println(reverse(120))         // 21
	fmt.Println(reverse(-123))        // -321
	fmt.Println(reverse(1534236469))  // 0 (przekracza zakres int32)
	fmt.Println(reverse(-2147483648)) // 0 (przekracza zakres)
}
