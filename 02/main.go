package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/giorgioprevitera/advent-of-code-2021/advent"
)

func getPartOneAnswer(input *[]string) int {
	h := 0
	d := 0
	for _, n := range *input {
		direction := strings.Split(n, " ")[0]
		value, _ := strconv.Atoi(strings.Split(n, " ")[1])
		if direction == "forward" {
			h += value
		}
		if direction == "down" {
			d += value
		}
		if direction == "up" {
			d -= value
		}
	}
	return (h * d)
}

func getPartTwoAnswer(input *[]string) int {
	h := 0
	d := 0
	a := 0
	for _, n := range *input {
		direction := strings.Split(n, " ")[0]
		value, _ := strconv.Atoi(strings.Split(n, " ")[1])
		if direction == "forward" {
			h += value
			d += value * a
		}
		if direction == "down" {
			a += value
		}
		if direction == "up" {
			a -= value
		}
	}
	return (h * d)
}

func main() {
	input, ok := advent.GetInput("input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	partOneAnswer := getPartOneAnswer(input)
	partTwoAnswer := getPartTwoAnswer(input)
	fmt.Printf("Part One answer is: %d\n", partOneAnswer)
	fmt.Printf("Part Two answer is: %d\n", partTwoAnswer)
}
