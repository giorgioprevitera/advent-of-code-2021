package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/giorgioprevitera/advent-of-code-2021/advent"
)

func getPartOneAnswer(input *[]string) int {
	gammaRate := ""
	epsilonRate := ""
	resultMap := countOnes(input)

	for _, n := range resultMap {
		if n >= len(*input)/2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	g, _ := strconv.ParseInt(gammaRate, 2, 0)
	e, _ := strconv.ParseInt(epsilonRate, 2, 0)
	return int(g) * int(e)
}

func countOnes(input *[]string) []int {
	resultMap := make([]int, len(strings.Split((*input)[0], "")))
	for _, row := range *input {
		for j, n := range row {
			if n == '1' {
				resultMap[j] += 1
			}
		}
	}
	return resultMap
}

func getPartTwoAnswer(input []string) int {
	oxygen := getRate(input, "oxygen", 0)
	co2 := getRate(input, "co2", 0)

	return oxygen * co2
}

func getRate(input []string, gasType string, index int) int {
	if len(input) == 1 {
		o, _ := strconv.ParseInt(input[0], 2, 0)
		return int(o)
	}

	var ones int
	var zeros int
	var r []string
	var valueToKeepWhenOnesAreGreater byte
	var valueToKeepWhenOnesAreLower byte

	switch gasType {
	case "oxygen":
		valueToKeepWhenOnesAreGreater = '1'
		valueToKeepWhenOnesAreLower = '0'
	case "co2":
		valueToKeepWhenOnesAreGreater = '0'
		valueToKeepWhenOnesAreLower = '1'

	}

	for _, row := range input {
		if row[index] == '1' {
			ones += 1
		} else {
			zeros += 1
		}
	}

	if ones >= zeros {
		for _, row := range input {
			if row[index] == valueToKeepWhenOnesAreGreater {
				r = append(r, row)
			}
		}
	} else {
		for _, row := range input {
			if row[index] == valueToKeepWhenOnesAreLower {
				r = append(r, row)
			}
		}

	}
	return getRate(r, gasType, index+1)
}

func main() {
	input := advent.GetStringInput("input.txt")
	partOneAnswer := getPartOneAnswer(input)
	fmt.Printf("Part One answer is: %d\n", partOneAnswer)
	partTwoAnswer := getPartTwoAnswer(*input)
	fmt.Printf("Part Two answer is: %d\n", partTwoAnswer)
}
