package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false

	fmt.Println(romanToInt("IV"))    //4
	fmt.Println(romanToInt("IX"))    //9
	fmt.Println(romanToInt("LVIII")) //58
}
