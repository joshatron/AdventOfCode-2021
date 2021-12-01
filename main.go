package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(day1Puzzle1())
	fmt.Println(day1Puzzle2())
}

func convertToInts(lines []string) []int {
	ints := make([]int, len(lines))

	for i, line := range lines {
		result, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting '", line, "' to an int")
		} else {
			ints[i] = result
		}
	}

	return ints
}

func splitIntoLines(content string) []string {
	return strings.Split(content, "\n")
}

func readDayInput(day string) string {
	content, err := ioutil.ReadFile("resources/day_" + day + ".txt")
	if err != nil {
		log.Fatal("Could not read content for day: " + day)
	}

	return string(content)
}

func day1Puzzle1() string {
	lines := convertToInts(splitIntoLines(readDayInput("01")))

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

func day1Puzzle2() string {
	lines := convertToInts(splitIntoLines(readDayInput("01")))

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
