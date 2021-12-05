package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestIsWinner(t *testing.T) {
	testCases := []struct {
		name  string
		board Board
		want  bool
	}{
		{"winnig board is true", Board{{1, 2, 3}, {-1, -1, -1}}, true},
		{"winnig board is true", Board{{-1, -1, -1}, {1, -1, -1}}, true},
		{"winnig board is true", Board{
			{-1, 1, 1},
			{-1, 1, 1},
			{-1, 1, 1},
			{-1, 1, 1},
		}, true},
		{"winnig board is true", Board{
			{-1, -1, -1},
			{-1, 1, 1},
			{-1, 1, 1},
			{-1, 1, 1},
		}, true},
		{"not winnig board is false", Board{{1, 2, 3}, {1, -1, -1}}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := tc.board.isWinner()
			assert.Equal(t, tc.want, actualResult)
		})
	}
}

/*
[
    [0, 1, 2, 3], 0
    [0, 1, 2, 3], 1
    [0, 1, 2, 3], 2
    [0, 1, 2, 3], 3
    [0, 1, 2, 3], 4
    [0, 1, 2, 3], 5
    [0, 1, 2, 3], 6
]
*/
