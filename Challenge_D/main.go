package main

import (
	"fmt"
	"math"
)

func getMaxMin(input []int) ([]float64, []float64) {
	max, max2, min, min2, maxSum, minSum := 0, 0, 0, 0, math.Inf(-1), math.Inf(1)
	for i := 1; i < len(input); i++ {
		if input[i] > input[max] {
			max = i
		}
		if input[i] < input[min] {
			min = i
		}
	}
	for i := 0; i < len(input); i++ {
		if i != max && float64(input[i]+input[max]) >= maxSum {
			max2 = i
			maxSum = float64(input[i] + input[max])
		}
		if i != min && float64(input[i]+input[min]) <= minSum {
			min2 = i
			minSum = float64(input[i] + input[min])
		}

	}
	minIndices := []float64{}
	maxIndices := []float64{}
	if min > min2 {
		minIndices = []float64{float64(min2), float64(min)}
	} else {
		minIndices = []float64{float64(min), float64(min2)}
	}
	if max > max2 {
		maxIndices = []float64{float64(max2), float64(max)}
	} else {
		maxIndices = []float64{float64(max), float64(max2)}
	}
	return minIndices, maxIndices
}

func main() {
	// Take length of string with len.
	input := []int{18, 2, 4, 5, 7}
	min, max := getMaxMin(input)
	fmt.Println("min = ", min)
	fmt.Println("max = ", max)
}
