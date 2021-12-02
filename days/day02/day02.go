package day02

import (
	"fmt"
	"strconv"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day2Puzzle1() string {
	lines := input.SplitIntoLines(input.ReadDayInput("02"))

	depth := 0
	position := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		magnitude, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			position += magnitude
		case "up":
			depth -= magnitude
		case "down":
			depth += magnitude
		default:
			fmt.Println("Unknown direction: ", direction)
		}
	}

	return fmt.Sprint(depth * position)
}

func Day2Puzzle2() string {
	lines := input.SplitIntoLines(input.ReadDayInput("02"))

	depth := 0
	position := 0
	aim := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		magnitude, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			position += magnitude
			depth += aim * magnitude
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		default:
			fmt.Println("Unknown direction: ", direction)
		}
	}

	return fmt.Sprint(depth * position)
}
