package days

import (
	"fmt"
	"strconv"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day05Puzzle1() string {
	segments := getSegments(input.SplitIntoLines(input.ReadDayInput("05")))
	pointsCovered := make(map[Point]int)

	for _, segment := range segments {
		if segment.isHorizontalOrVertical() {
			points := segment.getPoints()

			for _, point := range points {
				val, present := pointsCovered[point]
				if present {
					pointsCovered[point] = val + 1
				} else {
					pointsCovered[point] = 1
				}
			}
		}
	}

	intersections := 0
	for _, num := range pointsCovered {
		if num > 1 {
			intersections++
		}
	}

	return fmt.Sprint(intersections)
}

type Point struct {
	x, y int
}

type Segment struct {
	start, end Point
}

func (s Segment) isHorizontalOrVertical() bool {
	return s.start.x == s.end.x || s.start.y == s.end.y
}

func (s Segment) getPoints() []Point {
	points := []Point{}

	if s.start.x == s.end.x {
		start := intMin(s.start.y, s.end.y)
		end := intMax(s.start.y, s.end.y)
		for i := start; i <= end; i++ {
			points = append(points, Point{s.start.x, i})
		}
	} else if s.start.y == s.end.y {
		start := intMin(s.start.x, s.end.x)
		end := intMax(s.start.x, s.end.x)
		for i := start; i <= end; i++ {
			points = append(points, Point{i, s.start.y})
		}
	} else if s.start.x < s.end.x && s.start.y < s.end.y {
		for i := 0; i <= s.end.x-s.start.x; i++ {
			points = append(points, Point{s.start.x + i, s.start.y + i})
		}
	} else if s.start.x < s.end.x && s.start.y > s.end.y {
		for i := 0; i <= s.end.x-s.start.x; i++ {
			points = append(points, Point{s.start.x + i, s.start.y - i})
		}
	} else if s.start.x > s.end.x && s.start.y < s.end.y {
		for i := 0; i <= s.start.x-s.end.x; i++ {
			points = append(points, Point{s.start.x - i, s.start.y + i})
		}
	} else if s.start.x > s.end.x && s.start.y > s.end.y {
		for i := 0; i <= s.start.x-s.end.x; i++ {
			points = append(points, Point{s.start.x - i, s.start.y - i})
		}
	}

	return points
}

func intMin(first, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func intMax(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

func getSegments(lines []string) []Segment {
	segments := []Segment{}

	for _, line := range lines {
		parts := strings.Fields(line)
		segments = append(segments, Segment{parsePoint(parts[0]), parsePoint(parts[2])})
	}

	return segments
}

func parsePoint(point string) Point {
	parts := strings.Split(point, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Point{x, y}
}

func Day05Puzzle2() string {
	segments := getSegments(input.SplitIntoLines(input.ReadDayInput("05")))
	pointsCovered := make(map[Point]int)

	for _, segment := range segments {
		points := segment.getPoints()

		for _, point := range points {
			val, present := pointsCovered[point]
			if present {
				pointsCovered[point] = val + 1
			} else {
				pointsCovered[point] = 1
			}
		}
	}

	intersections := 0
	for _, num := range pointsCovered {
		if num > 1 {
			intersections++
		}
	}

	return fmt.Sprint(intersections)
}
