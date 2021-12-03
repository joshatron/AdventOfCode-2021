package days

import (
	"fmt"
	"strconv"

	"joshatron.io/aoc2021/input"
)

func Day3Puzzle1() string {
	lines := input.SplitIntoLines(input.ReadDayInput("03"))

	var gammaB, epsilonB string
	for i := 0; i < len(lines[0]); i++ {
		if isOneMostSignificantAtIndex(lines, i) {
			gammaB += "1"
			epsilonB += "0"
		} else {
			gammaB += "0"
			epsilonB += "1"
		}
	}

	gamma := binaryStringToInt(gammaB)
	epsilon := binaryStringToInt(epsilonB)

	return fmt.Sprint(gamma * epsilon)
}

func isOneMostSignificantAtIndex(lines []string, index int) bool {
	var ones, zeros int
	for _, line := range lines {
		switch line[index] {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}

	return ones >= zeros
}

func binaryStringToInt(num string) int {
	value, _ := strconv.ParseInt(num, 2, 32)

	return int(value)
}

func Day3Puzzle2() string {
	lines := input.SplitIntoLines(input.ReadDayInput("03"))

	oxygenGeneratorList := lines
	for i := 0; len(oxygenGeneratorList) > 1; i++ {
		oxygenGeneratorList = filterAtIndex(oxygenGeneratorList, i, true)
	}
	co2ScrubberList := lines
	for i := 0; len(co2ScrubberList) > 1; i++ {
		co2ScrubberList = filterAtIndex(co2ScrubberList, i, false)
	}

	oxygenGenerator := binaryStringToInt(oxygenGeneratorList[0])
	co2Scrubber := binaryStringToInt(co2ScrubberList[0])

	return fmt.Sprint(oxygenGenerator * co2Scrubber)
}

func filterAtIndex(lines []string, index int, mostSignificant bool) []string {
	oneMostSignificant := isOneMostSignificantAtIndex(lines, index)

	newLines := []string{}
	for _, line := range lines {
		isOne := line[index] == '1'
		if shouldInclude(isOne, oneMostSignificant, mostSignificant) {
			newLines = append(newLines, line)
		}
	}

	return newLines
}

func shouldInclude(isOne, oneMostSignificant, mostSignificant bool) bool {
	if mostSignificant {
		if (isOne && oneMostSignificant) || (!isOne && !oneMostSignificant) {
			return true
		} else {
			return false
		}
	} else {
		if (isOne && !oneMostSignificant) || (!isOne && oneMostSignificant) {
			return true
		} else {
			return false
		}
	}
}
