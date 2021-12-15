package days

import (
	"fmt"

	"joshatron.io/aoc2021/input"
)

func Day15Puzzle1() string {
	return fmt.Sprint(parseChitonCave(input.SplitIntoLines(input.ReadDayInput("15"))).minRisk())
}

func parseChitonCave(lines []string) ChitonCave {
	riskLevels := [][]int{}

	for _, line := range lines {
		row := []int{}
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		riskLevels = append(riskLevels, row)
	}

	return ChitonCave{riskLevels}
}

type ChitonCave struct {
	riskLevels [][]int
}

func (c ChitonCave) at(x, y int) int {
	if x < 0 || x >= len(c.riskLevels[0]) || y < 0 || y >= len(c.riskLevels) {
		return 1000000
	} else {
		return c.riskLevels[y][x]
	}
}

func (c ChitonCave) dest() (int, int) {
	maxX := len(c.riskLevels[0]) - 1
	maxY := len(c.riskLevels) - 1
	return maxX, maxY
}

func (c ChitonCave) minRisk() int {
	destX, destY := c.dest()
	open := SimplePriorityQueue{AStarLoc{0, 0, 1000000000, 0}, nil}
	open = pushToQueue(open, AStarLoc{0, 0, 0, (destX + destY)})
	closed := make(map[Point]int)

	i := 0
	for open.next != nil {
		var q AStarLoc
		open, q = open.pop()
		for _, child := range children(q, c) {
			if child.x == destX && child.y == destY {
				return child.g
			} else if shouldAddChild(child, open, closed) {
				open = pushToQueue(open, child)
			}
		}
		closed[Point{q.x, q.y}] = q.f()
		i++
	}

	return 0
}

func children(q AStarLoc, c ChitonCave) []AStarLoc {
	destX, destY := c.dest()
	children := []AStarLoc{}
	if q.x+1 <= destX {
		children = append(children, AStarLoc{q.x + 1, q.y, q.g + c.at(q.x+1, q.y), (destX - (q.x + 1)) + (destY - q.y)})
	}
	if q.x-1 >= 0 {
		children = append(children, AStarLoc{q.x - 1, q.y, q.g + c.at(q.x-1, q.y), (destX - (q.x - 1)) + (destY - q.y)})
	}
	if q.y+1 <= destY {
		children = append(children, AStarLoc{q.x, q.y + 1, q.g + c.at(q.x, q.y+1), (destX - q.x) + (destY - (q.y + 1))})
	}
	if q.y-1 >= 0 {
		children = append(children, AStarLoc{q.x, q.y - 1, q.g + c.at(q.x, q.y-1), (destX - q.x) + (destY - (q.y - 1))})
	}

	return children
}

func shouldAddChild(child AStarLoc, open SimplePriorityQueue, closed map[Point]int) bool {
	for open.next != nil && open.item.f() < child.f() {
		if open.item.x == child.x && open.item.y == child.y {
			return false
		}
		open = *open.next
	}
	_, ok := closed[Point{child.x, child.y}]
	if ok {
		return false
	}
	if child.x < 0 || child.y < 0 {
		return false
	}

	return true
}

type AStarLoc struct {
	x, y, g, h int
}

func (a AStarLoc) f() int {
	return a.g + a.h
}

func Day15Puzzle2() string {
	return fmt.Sprint(parseBigChitonCave(input.SplitIntoLines(input.ReadDayInput("15"))).minRisk())
}

func parseBigChitonCave(lines []string) ChitonCave {
	riskLevels := [][]int{}

	originalRows := [][]int{}
	for _, line := range lines {
		originalRow := []int{}
		for _, r := range line {
			originalRow = append(originalRow, int(r-'0'))
		}

		row := []int{}
		for i := 0; i < 5; i++ {
			for _, c := range originalRow {
				row = append(row, incRiskLevel(c, i))
			}
		}
		originalRows = append(originalRows, row)
	}

	for i := 0; i < 5; i++ {
		for _, row := range originalRows {
			finalRow := []int{}
			for _, c := range row {
				finalRow = append(finalRow, incRiskLevel(c, i))
			}
			riskLevels = append(riskLevels, finalRow)
		}
	}

	return ChitonCave{riskLevels}
}

func incRiskLevel(i, amount int) int {
	if i+amount <= 9 {
		return i + amount
	} else {
		return (i+amount)%10 + 1
	}
}

type SimplePriorityQueue struct {
	item AStarLoc
	next *SimplePriorityQueue
}

func pushToQueue(q SimplePriorityQueue, item AStarLoc) SimplePriorityQueue {
	if q.item.f() > item.f() {
		new := SimplePriorityQueue{item, &q}
		return new
	}

	current := &q
	for current.next != nil && current.next.item.f() < item.f() {
		current = current.next
	}

	new := SimplePriorityQueue{item, current.next}
	current.next = &new
	return q
}

func (q *SimplePriorityQueue) pop() (SimplePriorityQueue, AStarLoc) {
	return *q.next, q.item
}
