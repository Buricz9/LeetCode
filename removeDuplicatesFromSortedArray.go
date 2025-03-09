package main

import "fmt"

func removeDuplicates(nums []int) int {
	count, current := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			count = 0
			nums[current] = nums[i]
			current++
		} else {
			count++
			if count <= 1 {
				nums[current] = nums[i]
				current++
			}
		}
	}
	return current
}
func callRemoveDuplicates() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 4, 5}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4}))
}
