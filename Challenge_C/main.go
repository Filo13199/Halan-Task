package main

import (
	"fmt"
	"strconv"
)

func generatePalindrome(input int) string {
	palindrome := ""
	if input%2 == 1 {
		palindrome += "0"
	} else if input%2 == 0 {
		palindrome += "00"
	}
	for i := 1; len(palindrome) < input; i++ {
		palindrome = strconv.Itoa(i) + palindrome + strconv.Itoa(i)
	}
	return palindrome
}

func main() {
	// Take length of string with len.
	input := 7
	palindrome := generatePalindrome(input)
	fmt.Println(palindrome)
}
