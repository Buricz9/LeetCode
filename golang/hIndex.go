package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	if citations[0] == 0 {
		return 0
	}

	i := 0
	for _, c := range citations {
		if c > i {
			i++
		} else {
			return i
		}
	}
	return i
}

func callHIndex() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5})) //3
	fmt.Println(hIndex([]int{0, 0, 0, 0}))    //0
	fmt.Println(hIndex([]int{1, 3, 1}))       //1
	fmt.Println(hIndex([]int{11, 15}))        //2
}
