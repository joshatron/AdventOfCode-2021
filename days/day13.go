package days

import (
	"fmt"
	"strconv"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day13Puzzle1() string {
	lines := input.SplitIntoLines(input.ReadDayInput("13"))
	paper := parsePaper(lines)
	folds := parseFolds(lines)

	paper.fold(folds[0])

	return fmt.Sprint(paper.unique())
}

type Paper struct {
	points []Point
}

func parsePaper(lines []string) Paper {
	points := []Point{}

	for _, line := range lines {
		if len(line) == 0 {
			break
		} else {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			points = append(points, Point{x, y})
		}
	}

	return Paper{points}
}

type Fold struct {
	horizontal bool
	point      int
}

func parseFolds(lines []string) []Fold {
	folds := []Fold{}

	for _, line := range lines {
		if strings.HasPrefix(line, "fold along x") {
			amount, _ := strconv.Atoi(strings.Split(line, "=")[1])
			folds = append(folds, Fold{false, amount})
		} else if strings.HasPrefix(line, "fold along y") {
			amount, _ := strconv.Atoi(strings.Split(line, "=")[1])
			folds = append(folds, Fold{true, amount})
		}
	}

	return folds
}

func (p *Paper) height() int {
	max := 0
	for _, point := range p.points {
		max = intMax(max, point.y)
	}

	return max
}

func (p *Paper) width() int {
	max := 0
	for _, point := range p.points {
		max = intMax(max, point.x)
	}

	return max
}

func (p *Paper) fold(fold Fold) {
	for i, point := range p.points {
		if fold.horizontal {
			if point.y > fold.point {
				p.points[i] = Point{point.x, fold.point - (point.y - fold.point)}
			}
		} else {
			if point.x > fold.point {
				p.points[i] = Point{fold.point - (point.x - fold.point), point.y}
			}
		}
	}
}

func (p *Paper) unique() int {
	total := 0
	for i, point := range p.points {
		unique := true
		for j := i + 1; j < len(p.points); j++ {
			if point == p.points[j] {
				unique = false
				break
			}
		}

		if unique {
			total++
		}
	}

	return total
}

func Day13Puzzle2() string {
	lines := input.SplitIntoLines(input.ReadDayInput("13"))
	paper := parsePaper(lines)

	for _, fold := range parseFolds(lines) {
		paper.fold(fold)
	}

	return paper.toString()
}

func (p *Paper) toString() string {
	finalOutput := ""

	for y := 0; y <= p.height(); y++ {
		for x := 0; x <= p.width(); x++ {
			if p.contains(Point{x, y}) {
				finalOutput += "█"
			} else {
				finalOutput += " "
			}
		}
		finalOutput += "\n"
	}

	return finalOutput
}

func (p *Paper) contains(point Point) bool {
	for _, dot := range p.points {
		if dot == point {
			return true
		}
	}

	return false
}
