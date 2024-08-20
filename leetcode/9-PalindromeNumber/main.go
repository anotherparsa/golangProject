package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindromeFirst(121))
	fmt.Println(isPalindromeFirst(122))
	fmt.Println(isPalindromeSecond(121))
	fmt.Println(isPalindromeSecond(122))
}

func isPalindromeFirst(number int) bool {
	if number < 0 {
		return false
	} else {
		originalNumber := number
		reversedNumber := 0
		for number != 0 {
			remainder := number % 10
			number /= 10
			reversedNumber = reversedNumber*10 + remainder
		}
		return reversedNumber == originalNumber
	}
}

func isPalindromeSecond(number int) bool {
	stringNumber := strconv.Itoa(number)
	runes := []rune(stringNumber)
	originalText := string(stringNumber)
	i := 0
	j := len(runes) - 1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes) == originalText
}
