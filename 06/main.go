package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(filename string) []int {
	result := []int{}

	fileContent, _ := ioutil.ReadFile(filename)
	for _, c := range strings.Split(strings.TrimSpace(string(fileContent)), ",") {
		i, _ := strconv.Atoi(c)
		result = append(result, i)
	}

	return result
}

func countFish(input []int, days int) int {
	fishes := make(map[int]int, 8)
	result := 0

	for _, f := range input {
		fishes[f] += 1
	}
	for i := 0; i < days; i++ {
		newFishes := fishes[0]
		fishes[0] = fishes[1]
		fishes[1] = fishes[2]
		fishes[2] = fishes[3]
		fishes[3] = fishes[4]
		fishes[4] = fishes[5]
		fishes[5] = fishes[6]
		fishes[6] = fishes[7] + newFishes
		fishes[7] = fishes[8]
		fishes[8] = newFishes
	}
	for _, n := range fishes {
		result += n
	}
	return result
}

func main() {
	input := parseInput("input.txt")
	partOneAnswer := countFish(input, 80)
	fmt.Println("Answer to part one:", partOneAnswer)
	partTwoAnswer := countFish(input, 256)
	fmt.Println("Answer to part two:", partTwoAnswer)
}
