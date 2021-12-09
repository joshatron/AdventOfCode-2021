package days

import (
	"fmt"
	"math"

	"joshatron.io/aoc2021/input"
)

func Day09Puzzle1() string {
	heightMap := parseHeightMap(input.SplitIntoLines(input.ReadDayInput("09")))

	totalRisk := 0
	for x := 0; x < heightMap.xSize(); x++ {
		for y := 0; y < heightMap.ySize(); y++ {
			if heightMap.localMin(x, y) {
				totalRisk += heightMap.heightAt(x, y) + 1
			}
		}
	}

	return fmt.Sprint(totalRisk)
}

type HeightMap struct {
	heights [][]int
	visited [][]bool
}

func parseHeightMap(lines []string) HeightMap {
	heights := [][]int{}
	visited := [][]bool{}
	for _, line := range lines {
		heightRow := []int{}
		visitedRow := []bool{}
		for _, spot := range line {
			heightRow = append(heightRow, int(spot-'0'))
			visitedRow = append(visitedRow, false)
		}
		heights = append(heights, heightRow)
		visited = append(visited, visitedRow)
	}

	return HeightMap{heights, visited}
}

func (h HeightMap) xSize() int {
	return len(h.heights[0])
}

func (h HeightMap) ySize() int {
	return len(h.heights)
}

func (h HeightMap) heightAt(x, y int) int {
	if x >= 0 && x < h.xSize() && y >= 0 && y < h.ySize() {
		return h.heights[y][x]
	} else {
		return 10
	}
}

func (h HeightMap) hasVisited(x, y int) bool {
	if x >= 0 && x < h.xSize() && y >= 0 && y < h.ySize() {
		return h.visited[y][x]
	} else {
		return true
	}
}

func (h HeightMap) localMin(x, y int) bool {
	return h.heightAt(x, y) < h.heightAt(x+1, y) && h.heightAt(x, y) < h.heightAt(x-1, y) &&
		h.heightAt(x, y) < h.heightAt(x, y+1) && h.heightAt(x, y) < h.heightAt(x, y-1)
}

func Day09Puzzle2() string {
	heightMap := parseHeightMap(input.SplitIntoLines(input.ReadDayInput("09")))

	biggest := initBiggest(3)
	for x := 0; x < heightMap.xSize(); x++ {
		for y := 0; y < heightMap.ySize(); y++ {
			if !heightMap.hasVisited(x, y) && heightMap.heightAt(x, y) < 9 {
				biggest.add(basinSize(heightMap, x, y))
			}
		}
	}

	return fmt.Sprint(biggest.biggest[0] * biggest.biggest[1] * biggest.biggest[2])
}

func basinSize(heightMap HeightMap, x int, y int) int {
	total := 1

	heightMap.visited[y][x] = true
	if !heightMap.hasVisited(x+1, y) && heightMap.heightAt(x+1, y) < 9 {
		total += basinSize(heightMap, x+1, y)
	}
	if !heightMap.hasVisited(x-1, y) && heightMap.heightAt(x-1, y) < 9 {
		total += basinSize(heightMap, x-1, y)
	}
	if !heightMap.hasVisited(x, y+1) && heightMap.heightAt(x, y+1) < 9 {
		total += basinSize(heightMap, x, y+1)
	}
	if !heightMap.hasVisited(x, y-1) && heightMap.heightAt(x, y-1) < 9 {
		total += basinSize(heightMap, x, y-1)
	}

	return total
}

type BiggestX struct {
	biggest []int
}

func initBiggest(num int) BiggestX {
	nums := []int{}
	for i := 0; i < num; i++ {
		nums = append(nums, math.MinInt)
	}

	return BiggestX{nums}
}

func (b *BiggestX) add(n int) {
	for i, num := range b.biggest {
		if n > num {
			for j := len(b.biggest) - 2; j >= i; j-- {
				b.biggest[j+1] = b.biggest[j]
			}
			b.biggest[i] = n
			break
		}
	}
}
