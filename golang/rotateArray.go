package main

import "fmt"

func reverseNumTab(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k%n == 0 {
		return nums
	}

	k = k % n
	reverseNumTab(nums, 0, n-1)
	reverseNumTab(nums, 0, k-1)
	reverseNumTab(nums, k, n-1)
	return nums
}

func callRotateArray() {
	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)) // [5,6,7,1,2,3,4]
	fmt.Println(rotate([]int{-1, -100, 3, 99}, 2))     // [3,99,-1,-100]
}
