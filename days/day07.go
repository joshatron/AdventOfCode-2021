package days

import (
	"fmt"
	"math"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day07Puzzle1() string {
	positions := input.ConvertToInts(strings.Split(input.ReadDayInput("07"), ","))
	start, end := minAndMax(positions)

	minFuel := math.MaxInt
	for i := start; i <= end; i++ {
		fuel := fuelForPosition(positions, i)
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return fmt.Sprint(minFuel)
}

func minAndMax(ints []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt

	for _, i := range ints {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	return min, max
}

func fuelForPosition(positions []int, finalPosition int) int {
	fuel := 0
	for _, pos := range positions {
		fuel += intAbs(pos - finalPosition)
	}

	return fuel
}

func intAbs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func Day07Puzzle2() string {
	positions := input.ConvertToInts(strings.Split(input.ReadDayInput("07"), ","))
	start, end := minAndMax(positions)

	minFuel := math.MaxInt
	for i := start; i <= end; i++ {
		fuel := fuelForPosition2(positions, i)
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return fmt.Sprint(minFuel)
}

func fuelForPosition2(positions []int, finalPosition int) int {
	fuel := 0
	for _, pos := range positions {
		dist := intAbs(pos - finalPosition)
		fuel += (dist * (dist + 1)) / 2
	}

	return fuel
}
