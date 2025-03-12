package main

import (
	"fmt"
	"sort"
)

func prisonBreak(n int32, m int32, h []int32, v []int32) int64 {

	holeY := checkHole(v)
	holeX := checkHole(h)

	area := int64(holeX+1) * int64(holeY+1)

	return area
}

func checkHole(l []int32) int64 {
	length, maxLength := 1, 1

	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	prev := l[0]
	for i := 1; i < len(l); i++ {
		if prev+1 == l[i] {
			length++
		} else {
			length = 1
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return int64(maxLength)
}

func callPrisonBreak() {
	fmt.Println(prisonBreak(6, 6, []int32{4}, []int32{2}))
	fmt.Println(prisonBreak(6, 6, []int32{4, 5}, []int32{2, 5}))
}
