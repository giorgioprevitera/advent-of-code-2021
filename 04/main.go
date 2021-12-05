package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Numbers []int

type Board []Numbers

type Boards []Board

func remove(slice Numbers, s int) Numbers {
	return append(slice[:s], slice[s+1:]...)
}

func parseInput(filename string) (Numbers, Boards) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	numbers := Numbers{}
	boards := Boards{}

	var currentBoard Board
	var currentRow int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			currentBoard = make(Board, 5)
			currentRow = 0
			boards = append(boards, currentBoard)
			continue
		}

		if len(numbers) == 0 {
			s := strings.Split(line, ",")
			for _, v := range s {
				n, _ := strconv.Atoi(v)
				numbers = append(numbers, n)
			}
		} else {
			currentBoard[currentRow] = Numbers{}
			s := strings.Fields(line)
			for _, v := range s {
				n, _ := strconv.Atoi(v)
				currentBoard[currentRow] = append(currentBoard[currentRow], n)

			}
			currentRow += 1
		}
	}

	return numbers, boards
}

func (board Board) isWinner() bool {
	for _, row := range board {
		winner := true
		for _, n := range row {
			if n != -1 {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}

	for i := 0; i < len(board); i++ {
		winner := true
		for j := range board {
			if board[j][i] != -1 {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}

	return false

}

func (board Board) getResult() int {
	result := 0
	for _, row := range board {
		for _, n := range row {
			if n != -1 {
				result += n
			}
		}

	}
	return result

}

func getPartTwoAnswer(numbers Numbers, boards Boards) int {
	var winningBoard int
	var winningNumber int
	winners := map[int]bool{}

MainLoop:
	for _, drawnNumber := range numbers {
		for boardNumber, board := range boards {
			if _, ok := winners[boardNumber]; !ok {
				for rowNumber, row := range board {
					for columnNumber, number := range row {
						if number == drawnNumber {
							board[rowNumber][columnNumber] = -1
							if board.isWinner() {
								winners[boardNumber] = true
								if len(winners) == len(boards) {
									winningBoard = boardNumber
									winningNumber = drawnNumber
									break MainLoop
								}
							}
						}
					}
				}
			}
		}
	}

	result := boards[winningBoard].getResult()
	return result * winningNumber
}

func getPartOneAnswer(numbers Numbers, boards Boards) int {
	var winningBoard int
	var winningNumber int

MainLoop:
	for _, drawnNumber := range numbers {
		for boardNumber, board := range boards {
			for rowNumber, row := range board {
				for columnNumber, number := range row {
					if number == drawnNumber {
						board[rowNumber][columnNumber] = -1
					}
					if board.isWinner() {
						winningBoard = boardNumber
						winningNumber = drawnNumber
						break MainLoop
					}
				}
			}
		}
	}
	result := boards[winningBoard].getResult()
	return result * winningNumber
}

func main() {
	numbers, boards := parseInput("input.txt")

	partOneAnswer := getPartOneAnswer(numbers, boards)
	fmt.Printf("Part One answer is: %d\n", partOneAnswer)

	partTwoAnswer := getPartTwoAnswer(numbers, boards)
	fmt.Printf("Part Two answer is: %d\n", partTwoAnswer)

}
