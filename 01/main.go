package main

import (
	"fmt"
	"log"

	"github.com/giorgioprevitera/advent-of-code-2021/advent"
)

func sumSlice(input []int) (result int) {
	for _, n := range input {
		result += n
	}
	return
}

func getPartOneAnswer(input *[]int) int {
	count := 0
	for i, n := range *input {
		if i != 0 && n > (*input)[i-1] {
			count += 1
		}
	}
	return count
}

func getPartTwoAnswer(input *[]int) int {
	result := 0
	windowSize := 3
	for i := 0; i < len(*input)-windowSize; i++ {
		firstWindowSum := sumSlice((*input)[i : i+windowSize])
		secondWindowSum := sumSlice((*input)[i+1 : i+windowSize+1])
		if secondWindowSum > firstWindowSum {
			result += 1
		}
	}
	return result
}

func main() {
	input, ok := advent.GetInput("input.txt").(*[]int)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	partOneAnswer := getPartOneAnswer(input)
	partTwoAnswer := getPartTwoAnswer(input)
	fmt.Printf("Part One answer is: %d\n", partOneAnswer)
	fmt.Printf("Part Two answer is: %d\n", partTwoAnswer)
}
