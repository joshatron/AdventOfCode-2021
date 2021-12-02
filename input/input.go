package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ConvertToInts(lines []string) []int {
	ints := make([]int, len(lines))

	for i, line := range lines {
		result, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting '", line, "' to an int")
		} else {
			ints[i] = result
		}
	}

	return ints
}

func SplitIntoLines(content string) []string {
	return strings.Split(content, "\n")
}

func ReadDayInput(day string) string {
	content, err := ioutil.ReadFile("resources/day_" + day + ".txt")
	if err != nil {
		log.Fatal("Could not read content for day: " + day)
	}

	return string(content)
}
