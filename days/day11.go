package days

import (
	"fmt"
	"unicode"

	"joshatron.io/aoc2021/input"
)

func Day11Puzzle1() string {
	grid := parseOctoGrid(input.ReadDayInput("11"))

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += grid.step()
	}

	return fmt.Sprint(flashes)
}

type OctoGrid struct {
	octopi []int
	width  int
	height int
}

func parseOctoGrid(input string) OctoGrid {
	grid := []int{}
	width := 0
	height := 0
	counting := true
	for _, r := range input {
		if unicode.IsDigit(r) {
			grid = append(grid, int(r-'0'))
			if counting {
				width++
			}
		} else {
			counting = false
			height++
		}
	}
	height++

	return OctoGrid{grid, width, height}
}

func (o *OctoGrid) step() int {
	flashes := 0
	runAgain := true
	first := true
	for runAgain {
		runAgain = false
		for x := 0; x < o.width; x++ {
			for y := 0; y < o.height; y++ {
				if first {
					o.increment(x, y)
				}
				if o.at(x, y) > 9 && o.at(x, y) < 1000 {
					o.flash(x, y)
					o.set(x, y, 1000)
					flashes++
					runAgain = true
				}
			}
		}
		first = false
	}

	o.resetEnergyLevels()

	return flashes
}

func (o *OctoGrid) at(x, y int) int {
	if x >= 0 && x < o.width && y >= 0 && y < o.height {
		return o.octopi[y*o.width+x]
	} else {
		return 0
	}
}

func (o *OctoGrid) increment(x, y int) {
	if x >= 0 && x < o.width && y >= 0 && y < o.height {
		o.octopi[y*o.width+x]++
	}
}

func (o *OctoGrid) set(x, y, num int) {
	if x >= 0 && x < o.width && y >= 0 && y < o.height {
		o.octopi[y*o.width+x] = num
	}
}

func (o *OctoGrid) flash(x, y int) {
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			o.increment(x+dx, y+dy)
		}
	}
}

func (o *OctoGrid) resetEnergyLevels() {
	for x := 0; x < o.width; x++ {
		for y := 0; y < o.height; y++ {
			if o.at(x, y) > 9 {
				o.set(x, y, 0)
			}
		}
	}
}

func Day11Puzzle2() string {
	grid := parseOctoGrid(input.ReadDayInput("11"))

	step := 0
	for !grid.allFlashed() {
		grid.step()
		step++
	}

	return fmt.Sprint(step)
}

func (o *OctoGrid) allFlashed() bool {
	for x := 0; x < o.width; x++ {
		for y := 0; y < o.height; y++ {
			if o.at(x, y) != 0 {
				return false
			}
		}
	}

	return true
}
