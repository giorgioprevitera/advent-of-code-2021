package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Segment struct {
	start Coords
	end   Coords
}

type Segments []Segment

type Grid map[Coords]int

func (g Grid) count() int {
	result := 0
	for _, v := range g {
		if v > 1 {
			result += 1
		}
	}
	return result
}

func NewSegment(sx, ex, sy, ey int) Segment {
	s := Segment{}
	s.start.x = sx
	s.end.x = ex
	s.start.y = sy
	s.end.y = ey
	return s
}

func (s *Segment) sort() {
	if s.start.x > s.end.x {
		t := s.start.x
		s.start.x = s.end.x
		s.end.x = t
	}
	if s.start.y > s.end.y {
		t := s.start.y
		s.start.y = s.end.y
		s.end.y = t
	}

}

func parseInput(filename string) Segments {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var segments Segments

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		firstSet := strings.Fields(line)[0]
		secondSet := strings.Fields(line)[2]
		sx, _ := strconv.Atoi(strings.Split(firstSet, ",")[0])
		sy, _ := strconv.Atoi(strings.Split(firstSet, ",")[1])
		ex, _ := strconv.Atoi(strings.Split(secondSet, ",")[0])
		ey, _ := strconv.Atoi(strings.Split(secondSet, ",")[1])
		s := NewSegment(sx, ex, sy, ey)
		segments = append(segments, s)
	}
	return segments
}

func (s Segment) isValid() bool {
	return (s.start.x == s.end.x) || (s.start.y == s.end.y)
}

func getPartOneAnswer(input Segments) int {
	grid := Grid{}
	for _, segment := range input {
		segment.sort()
		if segment.isValid() {
			for i := segment.start.x; i <= segment.end.x; i++ {
				for j := segment.start.y; j <= segment.end.y; j++ {
					grid[Coords{i, j}] += 1
				}
			}
		}
	}
	return grid.count()
}

func getPartTwoAnswer(input Segments) int {
	grid := Grid{}
	for _, segment := range input {
		if segment.isValid() {
			segment.sort()
			for i := segment.start.x; i <= segment.end.x; i++ {
				for j := segment.start.y; j <= segment.end.y; j++ {
					grid[Coords{i, j}] += 1
				}
			}
		} else {
			if segment.start.x > segment.end.x {
				if segment.start.y > segment.end.y {
					for i := 0; i <= (segment.start.x - segment.end.x); i++ {
						grid[Coords{segment.start.x - i, segment.start.y - i}] += 1
					}
				} else {
					for i := 0; i <= (segment.start.x - segment.end.x); i++ {
						grid[Coords{segment.start.x - i, segment.start.y + i}] += 1
					}
				}
			} else {
				if segment.start.y > segment.end.y {
					for i := 0; i <= (segment.end.x - segment.start.x); i++ {
						grid[Coords{segment.start.x + i, segment.start.y - i}] += 1
					}
				} else {
					for i := 0; i <= (segment.end.x - segment.start.x); i++ {
						grid[Coords{segment.start.x + i, segment.start.y + i}] += 1
					}
				}
			}
		}
	}
	return grid.count()
}

func (g Grid) print(size int) {
	printableGrid := make([][]string, size)
	for v := range printableGrid {
		printableGrid[v] = []string{}
		for i := 0; i < size; i++ {
			printableGrid[v] = append(printableGrid[v], ".")

		}
	}

	for k, v := range g {
		printableGrid[k.y][k.x] = strconv.Itoa(v)

	}

	for v := range printableGrid {
		fmt.Println(printableGrid[v])
	}
}

func main() {
	input := parseInput("input.txt")

	partOneAnswer := getPartOneAnswer(input)
	fmt.Printf("Part One answer is: %d\n", partOneAnswer)
	partTwoAnswer := getPartTwoAnswer(input)
	fmt.Printf("Part Two answer is: %d\n", partTwoAnswer)
}
