package days

import (
	"fmt"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day06Puzzle1() string {
	fish := input.ConvertToInts(strings.Split(input.ReadDayInput("06"), ","))

	for i := 0; i < 80; i++ {
		fish = progressDay(fish)
	}

	return fmt.Sprint(len(fish))
}

func progressDay(fish []int) []int {
	newFish := []int{}

	for _, f := range fish {
		if f == 0 {
			newFish = append(newFish, 6)
			newFish = append(newFish, 8)
		} else {
			newFish = append(newFish, f-1)
		}
	}

	return newFish
}

func Day06Puzzle2() string {
	fish := input.ConvertToInts(strings.Split(input.ReadDayInput("06"), ","))
	counts := convertToCounts(fish)

	for i := 0; i < 256; i++ {
		counts = progressDayFast(counts)
	}

	return fmt.Sprint(getTotal(counts))
}

func convertToCounts(fish []int) []int {
	counts := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, f := range fish {
		counts[f]++
	}

	return counts
}

func progressDayFast(counts []int) []int {
	newCounts := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i, count := range counts {
		if i == 0 {
			newCounts[8] = count
			newCounts[6] = count
		} else {
			newCounts[i-1] += count
		}
	}

	return newCounts
}

func getTotal(counts []int) int {
	total := 0
	for _, count := range counts {
		total += count
	}

	return total
}
