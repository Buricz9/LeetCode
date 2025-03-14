package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {

	totalGas := 0
	currentGasLeft := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		currentGasLeft += gas[i] - cost[i]
		if currentGasLeft < 0 {
			start = i + 1
			currentGasLeft = 0
		}
	}

	if totalGas < 0 {
		return -1
	}
	return start
}

func callCanCompleteCircuit() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) //3
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))             //3

}
