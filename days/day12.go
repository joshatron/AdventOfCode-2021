package days

import (
	"fmt"
	"strings"
	"unicode"

	"joshatron.io/aoc2021/input"
)

func Day12Puzzle1() string {
	caves := parseCaves(input.SplitIntoLines(input.ReadDayInput("12")))
	return fmt.Sprint(len(getPaths(caves)))
}

type CaveSystem struct {
	caves map[string][]string
}

func parseCaves(connections []string) CaveSystem {
	allCaves := make(map[string][]string)

	for _, connection := range connections {
		sides := strings.Split(connection, "-")
		side1, ok1 := allCaves[sides[0]]
		if !ok1 {
			side1 = []string{}
		}
		side2, ok2 := allCaves[sides[1]]
		if !ok2 {
			side2 = []string{}
		}

		side1 = append(side1, sides[1])
		side2 = append(side2, sides[0])
		allCaves[sides[0]] = side1
		allCaves[sides[1]] = side2
	}

	return CaveSystem{allCaves}
}

type Path struct {
	path []string
}

func (p Path) contains(cave string) bool {
	for _, c := range p.path {
		if c == cave {
			return true
		}
	}

	return false
}

func getPaths(caves CaveSystem) []Path {
	return getPathsRecurse(caves, []Path{}, Path{[]string{"start"}})
}

func getPathsRecurse(caves CaveSystem, completed []Path, current Path) []Path {
	for _, option := range caves.caves[current.path[len(current.path)-1]] {
		if option == "end" {
			completed = append(completed, Path{append(current.path, option)})
		} else if (isSmall(option) && !current.contains(option)) || !isSmall(option) {
			completed = getPathsRecurse(caves, completed, Path{append(current.path, option)})
		}
	}

	return completed
}

func isSmall(cave string) bool {
	return unicode.IsLower([]rune(cave)[0])
}

func Day12Puzzle2() string {
	caves := parseCaves(input.SplitIntoLines(input.ReadDayInput("12")))
	return fmt.Sprint(len(getPaths2(caves)))
}

func getPaths2(caves CaveSystem) []Path {
	return getPathsRecurse2(caves, []Path{}, Path{[]string{"start"}}, false)
}

func getPathsRecurse2(caves CaveSystem, completed []Path, current Path, didSmall bool) []Path {
	for _, option := range caves.caves[current.path[len(current.path)-1]] {
		if option == "end" {
			completed = append(completed, Path{append(current.path, option)})
		} else if option == "start" {
			continue
		} else if isSmall(option) && current.contains(option) && !didSmall {
			completed = getPathsRecurse2(caves, completed, Path{append(current.path, option)}, true)
		} else if (isSmall(option) && !current.contains(option)) || !isSmall(option) {
			completed = getPathsRecurse2(caves, completed, Path{append(current.path, option)}, didSmall)
		}
	}

	return completed
}
