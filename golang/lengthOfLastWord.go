package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count != 0 && s[i] == ' ' {
			return count
		} else if s[i] != ' ' {
			count++
		}
	}

	return count
}

func callLengthOfLastWord() {
	fmt.Println(lengthOfLastWord("Hello World"))                 //5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) //4
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       //6
}
