package main

import "fmt"

func balancedSum(arr []int32) int32 {
	var totalSum int32
	for i := 0; i < len(arr); i++ {
		totalSum += arr[i]
	}

	var leftSum int32

	for i, val := range arr {
		rightSum := totalSum - leftSum - val
		if leftSum == rightSum {
			return int32(i)
		}
		leftSum += val
	}

	return -1

}

func callBalancedArray() {
	fmt.Println(balancedSum([]int32{1, 2, 3, 4, 6}))
}
