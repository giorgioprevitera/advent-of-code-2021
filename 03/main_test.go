package main

import (
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2021/advent"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := advent.GetStringInput("test_input.txt")
	actualOutput := getPartOneAnswer(input)
	expectedOutput := 198
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input := advent.GetStringInput("test_input.txt")
	actualOutput := getPartTwoAnswer(*input)
	expectedOutput := 230
	assert.Equal(t, expectedOutput, actualOutput)
}

func BenchmarkTest(b *testing.B) {
	input := advent.GetStringInput("test_input.txt")
	for i := 0; i < b.N; i++ {
		getPartTwoAnswer(*input)
	}
}
