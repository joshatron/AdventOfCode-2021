package days

import (
	"fmt"
	"math"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day14Puzzle1() string {
	lines := input.SplitIntoLines(input.ReadDayInput("14"))

	polymer := parseInitialPolymerFast(lines[0])
	insertions := parsePairInsertionsFast(lines[2:])

	for i := 0; i < 10; i++ {
		polymer = polymerizeFast(polymer, insertions)
	}

	runes := []rune(lines[0])
	return fmt.Sprint(mostCommonMinusLeastCommonElementCountsFast(polymer, runes[0], runes[len(runes)-1]))
}

func Day14Puzzle2() string {
	lines := input.SplitIntoLines(input.ReadDayInput("14"))

	polymer := parseInitialPolymerFast(lines[0])
	insertions := parsePairInsertionsFast(lines[2:])

	for i := 0; i < 40; i++ {
		polymer = polymerizeFast(polymer, insertions)
	}

	runes := []rune(lines[0])
	return fmt.Sprint(mostCommonMinusLeastCommonElementCountsFast(polymer, runes[0], runes[len(runes)-1]))
}

func parseInitialPolymerFast(line string) map[string]int {
	polymer := make(map[string]int)
	chars := []rune(line)

	for i := 0; i < len(chars)-1; i++ {
		combo := string(chars[i]) + string(chars[i+1])
		count, ok := polymer[combo]
		if !ok {
			count = 0
		}
		polymer[combo] = count + 1
	}

	return polymer
}

func parsePairInsertionsFast(lines []string) map[string][]string {
	insertions := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		firstParts := []rune(parts[0])
		insertions[parts[0]] = []string{string(firstParts[0]) + parts[1], parts[1] + string(firstParts[1])}
	}

	return insertions
}

func polymerizeFast(polymer map[string]int, insertions map[string][]string) map[string]int {
	newPolymer := make(map[string]int)

	for combo, count := range polymer {
		for _, toInsert := range insertions[combo] {
			oldCount, ok := newPolymer[toInsert]
			if !ok {
				oldCount = 0
			}
			newPolymer[toInsert] = oldCount + count
		}
	}

	return newPolymer
}

func mostCommonMinusLeastCommonElementCountsFast(polymer map[string]int, first rune, last rune) int {
	counts := make(map[rune]int)

	for combo, count := range polymer {
		for _, r := range combo {
			c, ok := counts[r]
			if !ok {
				c = 0
			}
			counts[r] = c + count
		}
	}

	totals := []int{}
	for r, value := range counts {
		value = value / 2
		if r == first || r == last {
			value++
		}
		totals = append(totals, value)
	}

	return intMaxArr(totals) - intMinArr(totals)
}

func intMaxArr(ints []int) int {
	max := math.MinInt
	for _, i := range ints {
		max = intMax(i, max)
	}

	return max
}

func intMinArr(ints []int) int {
	min := math.MaxInt
	for _, i := range ints {
		min = intMin(i, min)
	}

	return min
}
