package main

import "fmt"

func maxArea(height []int) int {

	maxAmount := 0
	for i, j := 0, len(height)-1; i < j; {
		currentAmount := min(height[i], height[j]) * (j - i)
		if currentAmount > maxAmount {
			maxAmount = currentAmount
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return maxAmount
}

func callMaxArea() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))   // Oczekiwany wynik: 49
	fmt.Println(maxArea([]int{1, 1}))                        // Oczekiwany wynik: 1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))               // Oczekiwany wynik: 16
	fmt.Println(maxArea([]int{1, 2, 1}))                     // Oczekiwany wynik: 2
	fmt.Println(maxArea([]int{2, 3, 10, 5, 7, 8, 9}))        // Oczekiwany wynik: 36
	fmt.Println(maxArea([]int{0, 0, 0, 0, 0}))               // Oczekiwany wynik: 0
	fmt.Println(maxArea([]int{1, 2, 3, 4, 5, 25, 24, 3, 4})) // Oczekiwany wynik: 24
}
