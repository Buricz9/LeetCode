package main

import "fmt"

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result += syms[i]
		}
	}

	return result
}

func callIntToRoman() {
	fmt.Println(intToRoman(3333)) // Oczekiwany wynik: "MMMCCCXXXIII"
	fmt.Println(intToRoman(9))    // Oczekiwany wynik: "IX"
	fmt.Println(intToRoman(10))   // Oczekiwany wynik: "X"
	fmt.Println(intToRoman(999))  // Oczekiwany wynik: "CMXCIX"
	fmt.Println(intToRoman(3999)) // Oczekiwany wynik: "MMMCMXCIX"
	fmt.Println(intToRoman(58))   // Oczekiwany wynik: "LVIII"
	fmt.Println(intToRoman(1444)) // Oczekiwany wynik: "MCDXLIV"
	fmt.Println(intToRoman(2024)) // Oczekiwany wynik: "MMXXIV"
	fmt.Println(intToRoman(1987)) // Oczekiwany wynik: "MCMLXXXVII"
	fmt.Println(intToRoman(2500)) // Oczekiwany wynik: "MMD"
}
