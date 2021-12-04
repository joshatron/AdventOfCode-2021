package days

import (
	"fmt"
	"strings"

	"joshatron.io/aoc2021/input"
)

func Day04Puzzle1() string {
	lines := input.SplitIntoLines(input.ReadDayInput("04"))

	calls := getCalls(lines)
	boards := getBoards(lines)

	for _, call := range calls {
		for _, board := range boards {
			board.mark(call)
			if board.hasWon() {
				return fmt.Sprint(board.getScore(call))
			}
		}
	}

	return ""
}

type Board struct {
	spots [][]int
	marks [][]bool
}

func (b *Board) mark(num int) {
	for i, row := range b.spots {
		for j, val := range row {
			if val == num {
				b.marks[i][j] = true
			}
		}
	}
}

func (b *Board) hasWon() bool {
	for _, row := range b.marks {
		full := true
		for _, spot := range row {
			if !spot {
				full = false
				break
			}
		}

		if full {
			return true
		}
	}

	for i := 0; i < len(b.marks); i++ {
		full := true
		for _, row := range b.marks {
			if !row[i] {
				full = false
				break
			}
		}

		if full {
			return true
		}
	}

	return false
}

func (b *Board) getScore(call int) int {
	sum := 0
	for i, row := range b.marks {
		for j, mark := range row {
			if !mark {
				sum += b.spots[i][j]
			}
		}
	}

	return sum * call
}

func getCalls(lines []string) []int {
	return input.ConvertToInts(strings.Split(lines[0], ","))
}

func getBoards(lines []string) []Board {
	boards := []Board{}
	currentBoard := [][]int{}

	for _, line := range lines[2:] {
		if len(line) > 0 {
			currentBoard = append(currentBoard, input.ConvertToInts(strings.Fields(line)))
		} else {
			boards = append(boards, Board{currentBoard, generateMarks(currentBoard)})
			currentBoard = [][]int{}
		}
	}

	return boards
}

func generateMarks(board [][]int) [][]bool {
	marks := [][]bool{}
	for i := 0; i < len(board); i++ {
		markRow := []bool{}
		for j := 0; j < len(board[0]); j++ {
			markRow = append(markRow, false)
		}
		marks = append(marks, markRow)
	}

	return marks
}

func Day04Puzzle2() string {
	lines := input.SplitIntoLines(input.ReadDayInput("04"))

	calls := getCalls(lines)
	boards := getBoards(lines)

	toRemove := []int{}
	for _, call := range calls {
		for i, board := range boards {
			board.mark(call)
			if board.hasWon() {
				if len(boards) > 1 {
					toRemove = append(toRemove, i)
				} else {
					return fmt.Sprint(board.getScore(call))
				}
			}
		}

		boards = removeSelected(boards, toRemove)
		toRemove = []int{}
	}

	return ""
}

func removeSelected(boards []Board, toRemove []int) []Board {
	newBoards := []Board{}

	for i, board := range boards {
		found := false
		for _, num := range toRemove {
			if num == i {
				found = true
				break
			}
		}

		if !found {
			newBoards = append(newBoards, board)
		}
	}

	return newBoards
}
