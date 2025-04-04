package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix := make([]strings.Builder, numRows)
	count, check := 0, true

	for i := 0; i < len(s); i++ {
		matrix[count].WriteByte(s[i])
		if count == numRows-1 {
			check = false
		} else if count == 0 {
			check = true
		}

		if check {
			count++
		} else {
			count--
		}
	}

	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(matrix[i].String())
	}

	return result.String()
}

func callConversion() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
