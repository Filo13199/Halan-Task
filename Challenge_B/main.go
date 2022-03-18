package main

import "fmt"

func isAnagram(input1 string, input2 string) bool {
	input1Map := make(map[rune]int)
	input2Map := make(map[rune]int)
	//generating map for each respective input, using rune instead of byte to handle unicode characters
	for i := 0; i < len(input1); i++ {
		input1Map[rune(input1[i])] += 1
	}
	for i := 0; i < len(input2); i++ {
		input2Map[rune(input1[i])] += 1
	}
	//compare input1 map against input2 map
	for k := range input1Map {
		if input1Map[k] != input2Map[k] {
			return false
		}
	}
	return true
}

func main() {
	// Take length of string with len.
	input1 := "restful"
	input2 := "fluster"
	needleIndex := isAnagram(input1, input2)
	fmt.Println(needleIndex)
}
