package input

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func SeparateBySpaces(lines []string) [][]string {
	separated := make([][]string, len(lines))

	for i, line := range lines {
		separated[i] = strings.Split(line, " ")
	}

	return separated

}

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
