package days

import (
	"fmt"
	"sort"

	"joshatron.io/aoc2021/input"
)

func Day10Puzzle1() string {
	lines := input.SplitIntoLines(input.ReadDayInput("10"))

	total := 0
	for _, line := range lines {
		total += corruptScore(line)
	}

	return fmt.Sprint(total)
}

func corruptScore(line string) int {
	stack := initStack()

	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack.push(c)
		} else {
			toClose := stack.pop()
			if !bracketsMatch(toClose, c) {
				switch c {
				case ')':
					return 3
				case ']':
					return 57
				case '}':
					return 1197
				case '>':
					return 25137
				}
			}
		}
	}

	return 0
}

type Element struct {
	item rune
	prev *Element
}

type Stack struct {
	top Element
}

func initStack() Stack {
	return Stack{Element{0, nil}}
}

func (s *Stack) push(item rune) {
	oldTop := s.top
	s.top = Element{item, &oldTop}
}

func (s *Stack) pop() rune {
	if s.top.prev != nil {
		topRune := s.top.item
		s.top = *s.top.prev
		return topRune
	} else {
		return 0
	}
}

func bracketsMatch(open, close rune) bool {
	return (open == '(' && close == ')') || (open == '[' && close == ']') || (open == '{' && close == '}') || (open == '<' && close == '>')
}

func Day10Puzzle2() string {
	lines := input.SplitIntoLines(input.ReadDayInput("10"))

	scores := []int{}
	for _, line := range lines {
		score := autocompleteScore(line)
		if score > 0 {
			scores = append(scores, score)
		}
	}

	return fmt.Sprint(getMiddle(scores))
}

func autocompleteScore(line string) int {
	stack := initStack()

	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack.push(c)
		} else {
			toClose := stack.pop()
			if !bracketsMatch(toClose, c) {
				return 0
			}
		}
	}

	score := 0
	for stack.top.prev != nil {
		score *= 5
		switch stack.pop() {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}

	return score
}

func getMiddle(scores []int) int {
	sort.Ints(scores)
	return scores[len(scores)/2]
}
