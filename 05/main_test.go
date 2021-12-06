package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestPartOne(t *testing.T) {
	input := parseInput("test_input.txt")
	actualOutput := getPartOneAnswer(input)
	expectedOutput := 5
	assert.Equal(t, actualOutput, expectedOutput)
}

func TestPartTwo(t *testing.T) {
	input := parseInput("test_input.txt")
	actualOutput := getPartTwoAnswer(input)
	expectedOutput := 12
	assert.Equal(t, actualOutput, expectedOutput)
}
