package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for index, value := range nums {
		want := target - value

		if val, ok := seen[want]; ok {
			return []int{val, index}
		}

		seen[value] = index
	}

	return []int{}
}
