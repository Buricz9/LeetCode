package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	currentJump := nums[0]

	for i := 1; i < len(nums); i++ {
		currentJump--

		if currentJump < 0 {
			return false
		}

		if nums[i] > currentJump {
			currentJump = nums[i]
		}
	}
	return true
}

func callCanJump() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) //true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) //false
}
