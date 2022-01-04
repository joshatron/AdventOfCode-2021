package days

import (
	"fmt"
	"unicode"

	"joshatron.io/aoc2021/input"
)

func Day18Puzzle1() string {
	numbers := parseSnailfishNumbers(input.SplitIntoLines(input.ReadDayInput("18")))

	total := numbers[0]
	for _, number := range numbers[1:] {
		fmt.Println(toStringSnailfish(total))
		fmt.Println("+", toStringSnailfish(number))
		total = reduce(SnailfishNumber{-100, &total, -100, &number})
		fmt.Println("=", toStringSnailfish(total))
		fmt.Println()
	}
	fmt.Println("------------------------------------------------------------------")

	fmt.Println("[[[[3,0],[5,3]],[4,4]],[5,5]]")
	fmt.Println(toStringSnailfish(total))
	return fmt.Sprint(magnitude(total))
}

func parseSnailfishNumbers(lines []string) []SnailfishNumber {
	numbers := []SnailfishNumber{}
	for _, line := range lines {
		numbers = append(numbers, parseSnailfishNumber(line))
	}

	return numbers
}

func parseSnailfishNumber(input string) SnailfishNumber {
	number, _ := parseSnailfishNumberRecurse(input)

	return number
}

func parseSnailfishNumberRecurse(input string) (SnailfishNumber, string) {
	number := SnailfishNumber{}
	runes := []rune(input)
	if unicode.IsDigit(runes[1]) {
		number.xLiteral = int(runes[1] - '0')
		if unicode.IsDigit(runes[3]) {
			number.yLiteral = int(runes[3] - '0')
			return number, input[5:]
		} else {
			number.yLiteral = -100
			innerY, remaining := parseSnailfishNumberRecurse(input[3:])
			number.ySnailfishNumber = &innerY
			return number, remaining[1:]
		}
	} else {
		number.xLiteral = -100
		innerX, remaining := parseSnailfishNumberRecurse(input[1:])
		number.xSnailfishNumber = &innerX
		runes = []rune(remaining)
		if unicode.IsDigit(runes[1]) {
			number.yLiteral = int(runes[1] - '0')
			return number, remaining[3:]
		} else {
			number.yLiteral = -100
			innerY, remainingRight := parseSnailfishNumberRecurse(remaining[1:])
			number.ySnailfishNumber = &innerY
			return number, remainingRight[1:]
		}
	}
}

type SnailfishNumber struct {
	xLiteral         int
	xSnailfishNumber *SnailfishNumber
	yLiteral         int
	ySnailfishNumber *SnailfishNumber
}

func reduce(num SnailfishNumber) SnailfishNumber {
	stillGoing := true
	for stillGoing {
		var exploded, reduced bool
		num, exploded, _, _, reduced = reduceRecursive(num, 1)
		stillGoing = exploded || reduced
		fmt.Println(toStringSnailfish(num))
	}

	return num
}

func reduceRecursive(num SnailfishNumber, depth int) (SnailfishNumber, bool, int, int, bool) {
	if depth == 4 && num.xLiteral == -100 {
		left := num.xSnailfishNumber.xLiteral
		right := num.xSnailfishNumber.yLiteral
		num.xLiteral = 0
		num.xSnailfishNumber = nil
		if num.yLiteral != -100 {
			num.yLiteral += right
			right = -1
		}
		return num, true, left, right, false
	} else if num.xLiteral > 9 {
		halfMinus := num.xLiteral / 2
		num.xSnailfishNumber = &SnailfishNumber{halfMinus, nil, num.xLiteral - halfMinus, nil}
		num.xLiteral = -100
		return num, false, -1, -1, true
	} else if num.xLiteral == -100 {
		newNum, exploding, left, right, reducing := reduceRecursive(*num.xSnailfishNumber, depth+1)
		num.xSnailfishNumber = &newNum
		if exploding && right != -1 && num.yLiteral != -100 {
			num.yLiteral += right
			return num, true, left, -1, false
		} else if exploding || reducing {
			return num, exploding, left, right, reducing
		} else if num.yLiteral == -100 {
			newNumR, explodingR, leftR, rightR, reducingR := reduceRecursive(*num.ySnailfishNumber, depth+1)
			num.ySnailfishNumber = &newNumR
			return num, explodingR, leftR, rightR, reducingR
		} else {
			return num, exploding, left, right, reducing
		}
	} else if depth == 4 && num.yLiteral == -100 {
		right := num.ySnailfishNumber.yLiteral
		num.yLiteral = 0
		num.ySnailfishNumber = nil
		num.xLiteral += right
		return num, true, -1, right, false
	} else if num.yLiteral > 9 {
		halfMinus := num.yLiteral / 2
		num.ySnailfishNumber = &SnailfishNumber{halfMinus, nil, num.yLiteral - halfMinus, nil}
		num.yLiteral = -100
		return num, false, -1, -1, true
	} else if num.yLiteral == -100 {
		newNum, exploding, left, right, reducing := reduceRecursive(*num.ySnailfishNumber, depth+1)
		num.ySnailfishNumber = &newNum
		if exploding && left != -1 {
			num.xLiteral += left
			left = -1
		}
		return num, exploding, left, right, reducing
	} else {
		return num, false, -1, -1, false
	}
}

func magnitude(num SnailfishNumber) int {
	total := 0
	if num.xLiteral == -100 {
		total += 3 * magnitude(*num.xSnailfishNumber)
	} else {
		total += 3 * num.xLiteral
	}
	if num.yLiteral == -100 {
		total += 2 * magnitude(*num.ySnailfishNumber)
	} else {
		total += 2 * num.yLiteral
	}

	return total
}

func toStringSnailfish(num SnailfishNumber) string {
	str := "["
	if num.xLiteral == -100 {
		str += toStringSnailfish(*num.xSnailfishNumber)
	} else {
		str += fmt.Sprint(num.xLiteral)
	}
	str += ","
	if num.yLiteral == -100 {
		str += toStringSnailfish(*num.ySnailfishNumber)
	} else {
		str += fmt.Sprint(num.yLiteral)
	}
	str += "]"

	return str
}

func Day18Puzzle2() string {
	return ""
}
