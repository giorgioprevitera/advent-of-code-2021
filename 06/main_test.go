package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestPartOne(t *testing.T) {
	input := parseInput("test_input.txt")
	actualOutput := countFish(input, 80)
	expectedOutput := 5934
	assert.Equal(t, actualOutput, expectedOutput)
}

func TestPartTwo(t *testing.T) {
	input := parseInput("test_input.txt")
	actualOutput := countFish(input, 256)
	expectedOutput := 26984457539
	assert.Equal(t, actualOutput, expectedOutput)
}
