package main

import (
	"sort"
)

func frequencySort(items []int) []int {
	freq := make(map[int]int)
	for _, item := range items {
		freq[item]++
	}

	sort.Slice(items, func(i, j int) bool {
		if freq[items[i]] != freq[items[j]] {
			result := freq[items[i]] < freq[items[j]]
			return result
		}
		result := items[i] < items[j]

		return result
	})

	return items
}

func callFrequencySorted() {
	frequencySort([]int{1, 1, 2, 3, 4, 5})
}
