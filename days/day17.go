package days

import (
	"fmt"
	"strconv"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day17Puzzle1() string {
	target := parseTargetArea(input.ReadDayInput("17"))

	maxY := 0
	for i := 0; i < 10000; i++ {
		maxY = intMax(simulateY(i, target.startY, target.endY), maxY)
	}

	return fmt.Sprint(maxY)
}

func simulateY(initialV, startY, endY int) int {
	currentV := initialV
	yPos := 0
	maxY := 0
	for yPos > startY {
		yPos += currentV
		maxY = intMax(yPos, maxY)
		currentV--
		if yPos <= endY && yPos >= startY {
			return maxY
		}
	}

	return 0
}

func parseTargetArea(input string) TargetArea {
	sections := strings.Split(input, " ")
	startX, endX := extractNums(sections[2])
	startY, endY := extractNums(sections[3])

	return TargetArea{startX, endX, startY, endY}
}

func extractNums(section string) (int, int) {
	justNumbers := section[2:]
	if justNumbers[len(justNumbers)-1] == ',' {
		justNumbers = justNumbers[:len(justNumbers)-1]
	}
	splitNumbers := strings.Split(justNumbers, "..")
	start, _ := strconv.Atoi(splitNumbers[0])
	end, _ := strconv.Atoi(splitNumbers[1])

	return start, end
}

type TargetArea struct {
	startX, endX, startY, endY int
}

func Day17Puzzle2() string {
	target := parseTargetArea(input.ReadDayInput("17"))
	minX, maxX, minY, maxY := getLimits(target)

	count := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if hitsTargetArea(x, y, target) {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}

func hitsTargetArea(initialDX int, initialDY int, target TargetArea) bool {
	currentX := 0
	currentY := 0
	currentDX := initialDX
	currentDY := initialDY

	for currentX <= target.endX && currentY >= target.startY {
		currentX += currentDX
		currentY += currentDY
		currentDX = intMax(0, currentDX-1)
		currentDY--
		if currentX >= target.startX && currentX <= target.endX && currentY >= target.startY && currentY <= target.endY {
			return true
		}
	}

	return false
}

func getLimits(target TargetArea) (int, int, int, int) {
	minX := 0
	maxX := target.endX
	minY := target.startY
	maxY := 0

	for i := 0; i < 10000; i++ {
		if simulateY(i, target.startY, target.endY) > 0 {
			maxY = i
		}
	}

	for i := 0; i <= target.endX; i++ {
		if simulateX(i, target.startX, target.endX) > 0 && minX == 0 {
			minX = i
			break
		}
	}

	return minX, maxX, minY, maxY
}

func simulateX(initialV, startX, endX int) int {
	currentV := initialV
	currentX := 0

	for currentX <= endX && currentV > 0 {
		currentX += currentV
		currentV--
		if currentX >= startX && currentX <= endX {
			return currentX
		}
	}

	return 0
}
