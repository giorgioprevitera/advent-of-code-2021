package main

import (
	"log"
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2021/advent"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, ok := advent.GetInput("test_input.txt").(*[]int)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	actualOutput := getPartOneAnswer(input)
	expectedOutput := 7
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input, ok := advent.GetInput("test_input.txt").(*[]int)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	actualOutput := getPartTwoAnswer(input)
	expectedOutput := 5
	assert.Equal(t, expectedOutput, actualOutput)
}
