package main

import "fmt"

func main() {
	fmt.Println(RomanToInteger("IV"))
	fmt.Println(RomanToInteger("IX"))
	fmt.Println(RomanToInteger("VIII"))
}

func RomanToInteger(romanNumber string) int {
	TranslationMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0

	for i := 0; i < len(romanNumber); i++ {
		currentvalue := TranslationMap[romanNumber[i]]
		if i < len(romanNumber)-1 && currentvalue < TranslationMap[romanNumber[i+1]] {
			result -= currentvalue
		} else {
			result += currentvalue
		}
	}
	return result
}
