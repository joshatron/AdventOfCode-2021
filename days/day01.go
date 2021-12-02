package days

import (
	"fmt"
	"math"

	"joshatron.io/aoc2021/input"
)

func Day1Puzzle1() string {
	lines := input.ConvertToInts(input.SplitIntoLines(input.ReadDayInput("01")))

	previous := math.MaxInt
	increases := 0
	for _, depth := range lines {
		if depth > previous {
			increases++
		}
		previous = depth
	}

	return fmt.Sprint(increases)
}

func Day1Puzzle2() string {
	lines := input.ConvertToInts(input.SplitIntoLines(input.ReadDayInput("01")))

	increases := 0
	for i, depth := range lines {
		if i+3 < len(lines) {
			if depth < lines[i+3] {
				increases++
			}
		}
	}

	return fmt.Sprint(increases)
}
