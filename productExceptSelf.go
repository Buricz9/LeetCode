package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]

	}
	return answer
}

func callProductExceptSelf() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      //[24,12,8,6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) //[0,0,9,0,0]

}
