package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		str = strings.ToLower(str)
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	input := []string{
		"Radar",
		"Malam",
		"Kasur ini rusak",
		"Ibu Ratna antar ubi",
		"Malas",
		"Makan nasi goreng",
		"Balonku ada lima",
	}

	for i := 0; i < len(input); i++ {

		result := isPalindrome(input[i])

		if result {
			fmt.Printf("%s => palindrome\n", input[i])

		} else {
			fmt.Printf("%s => not palindrome\n", input[i])
		}
	}
}
