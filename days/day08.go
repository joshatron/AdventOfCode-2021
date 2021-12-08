package days

import (
	"fmt"
	"strconv"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day08Puzzle1() string {
	displays := getDisplays(input.SplitIntoLines(input.ReadDayInput("08")))

	count := 0
	for _, display := range displays {
		for _, output := range display.output {
			length := len(output)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}

	return fmt.Sprint(count)
}

type Display struct {
	patterns, output []string
}

func getDisplays(lines []string) []Display {
	displays := []Display{}

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		d := Display{strings.Fields(parts[0]), strings.Fields(parts[1])}
		d.unscramble()
		displays = append(displays, d)
	}

	return displays
}

func Day08Puzzle2() string {
	displays := getDisplays(input.SplitIntoLines(input.ReadDayInput("08")))

	total := 0
	for _, display := range displays {
		total += display.parseOutput()
	}

	return fmt.Sprint(total)
}

//0- length 6, shares 2 letters with 1 and 3 letters with 4
//1- length 2
//2- length 5, shares 1 letters with 1 and 2 letters with 4
//3- legnth 5, shares 2 letters with 1 and 3 letters with 4
//4- length 4
//5- length 5, shares 1 letters with 1 and 3 letters with 4
//6- length 6, shares 1 letters with 1 and 3 letters with 4
//7- length 3
//8- length 7
//9- length 6, shares 2 letters with 1 and 4 letters with 4
func (d *Display) unscramble() {
	unscrambled := []string{"", "", "", "", "", "", "", "", "", ""}

	//Initial pass to get the base ones done
	for _, num := range d.patterns {
		length := len(num)
		if length == 2 {
			unscrambled[1] = num
		} else if length == 4 {
			unscrambled[4] = num
		} else if length == 3 {
			unscrambled[7] = num
		} else if length == 7 {
			unscrambled[8] = num
		}
	}

	//Now for the remaining numbers
	for _, num := range d.patterns {
		length := len(num)
		oneCompare := charsShared(num, unscrambled[1])
		fourCompare := charsShared(num, unscrambled[4])
		if length == 6 && oneCompare == 2 && fourCompare == 3 {
			unscrambled[0] = num
		} else if length == 5 && oneCompare == 1 && fourCompare == 2 {
			unscrambled[2] = num
		} else if length == 5 && oneCompare == 2 && fourCompare == 3 {
			unscrambled[3] = num
		} else if length == 5 && oneCompare == 1 && fourCompare == 3 {
			unscrambled[5] = num
		} else if length == 6 && oneCompare == 1 && fourCompare == 3 {
			unscrambled[6] = num
		} else if length == 6 && oneCompare == 2 && fourCompare == 4 {
			unscrambled[9] = num
		}
	}

	d.patterns = unscrambled
}

func charsShared(first, second string) int {
	shared := 0
	for _, fc := range first {
		for _, sc := range second {
			if fc == sc {
				shared++
				break
			}
		}
	}

	return shared
}

func (d *Display) parseOutput() int {
	number := ""

	for _, num := range d.output {
		number += d.getNum(num)
	}

	finalNum, _ := strconv.Atoi(number)

	return finalNum
}

func (d *Display) getNum(num string) string {
	for i, pattern := range d.patterns {
		if len(num) == len(pattern) && len(num) == charsShared(num, pattern) {
			return fmt.Sprint(i)
		}
	}

	return ""
}
