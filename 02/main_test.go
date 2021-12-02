package main

import (
	"log"
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2021/advent"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, ok := advent.GetInput("test_input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	actualOutput := getPartOneAnswer(input)
	expectedOutput := 150
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input, ok := advent.GetInput("test_input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	actualOutput := getPartTwoAnswer(input)
	expectedOutput := 900
	assert.Equal(t, expectedOutput, actualOutput)
}
