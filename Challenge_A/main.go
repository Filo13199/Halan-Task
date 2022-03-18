package main

import "fmt"

func getFirstOccurenceOfNeedle(hayStack string, needle string) int {
	needleCounter := 0
	for i := 0; i < len(hayStack); i++ {
		if hayStack[i] == needle[needleCounter] {
			needleCounter++
			if needleCounter == len(needle) {
				return i - needleCounter + 1
			}
		} else {
			needleCounter = 0
		}

	}
	return -1
}

func main() {
	// Take length of string with len.
	hayStack := "consciousness”"
	needle := "cious"
	needleIndex := getFirstOccurenceOfNeedle(hayStack, needle)
	fmt.Println(needleIndex)
}
