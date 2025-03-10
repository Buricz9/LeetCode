package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev1 := 2 // (n-1)
	prev2 := 1 // (n-2)

	for i := 3; i <= n; i++ {
		current := prev2 + prev1
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func callClimbingStairs() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(5))
}
