package day4

import (
	"github.com/flyrell/aoc2024/internal/utils"
)

func Run() int {
	sum := 0
	word := []rune{'X', 'M', 'A', 'S'}
	xPositions := make([][]int, 0, 20)
	positionsMap := make(map[int]map[int]rune)
	utils.ReadInputFileByLine("/app/internal/day4/input", func(rowNo int, text string) {
		for colNo, char := range text {
			if char == word[0] {
				xPositions = append(xPositions, []int{rowNo, colNo})
				continue
			}
			if char != word[1] && char != word[2] && char != word[3] {
				continue
			}
			if _, ok := positionsMap[rowNo]; !ok {
				positionsMap[rowNo] = make(map[int]rune)
			}
			positionsMap[rowNo][colNo] = char
		}
	})

	for _, coords := range xPositions {
		x := coords[0]
		y := coords[1]

		if v, ok := positionsMap[x]; ok && isHorizontalWord(word[1:], y, &v) {
			sum++
		}

		if v, ok := positionsMap[x]; ok && isReversedHorizontalWord(word[1:], y, &v) {
			sum++
		}

		if isVerticalWord(word[1:], x, y, &positionsMap) {
			sum++
		}

		if isReversedVerticalWord(word[1:], x, y, &positionsMap) {
			sum++
		}

		if isDiagonalBottomRight(word[1:], x, y, &positionsMap, 1) {
			sum++
		}

		if isDiagonalBottomLeft(word[1:], x, y, &positionsMap, 1) {
			sum++
		}

		if isDiagonalTopRight(word[1:], x, y, &positionsMap, 1) {
			sum++
		}

		if isDiagonalTopLeft(word[1:], x, y, &positionsMap, 1) {
			sum++
		}
	}

	defer func() {
		println("Result of Day 4 is:", sum)
	}()

	return sum
}

func isHorizontalWord(word []rune, y int, positionsMap *map[int]rune) bool {
	for i, char := range word {
		if ch, ok := (*positionsMap)[y+i+1]; !ok || ch != char {
			return false
		}
	}
	return true
}

func isReversedHorizontalWord(word []rune, y int, positionsMap *map[int]rune) bool {
	for i, char := range word {
		if ch, ok := (*positionsMap)[y-i-1]; !ok || ch != char {
			return false
		}
	}
	return true
}

func isVerticalWord(word []rune, x int, y int, positionsMap *map[int]map[int]rune) bool {
	for i, char := range word {
		row, ok := (*positionsMap)[x+i+1]
		if !ok {
			return false
		}

		if ch, ok := row[y]; !ok || ch != char {
			return false
		}
	}
	return true
}

func isReversedVerticalWord(word []rune, x int, y int, positionsMap *map[int]map[int]rune) bool {
	for i, char := range word {
		row, ok := (*positionsMap)[x-i-1]
		if !ok {
			return false
		}

		if ch, ok := row[y]; !ok || ch != char {
			return false
		}
	}
	return true
}

func isDiagonalBottomRight(word []rune, x int, y int, positionsMap *map[int]map[int]rune, diff int) bool {
	for i, char := range word {
		row, ok := (*positionsMap)[x+i+diff]
		if !ok {
			return false
		}

		if ch, ok := row[y+i+diff]; !ok || ch != char {
			return false
		}
	}
	return true
}

func isDiagonalTopRight(word []rune, x int, y int, positionsMap *map[int]map[int]rune, diff int) bool {
	for i, char := range word {
		row, ok := (*positionsMap)[x-i-diff]
		if !ok {
			return false
		}

		if ch, ok := row[y+i+diff]; !ok || ch != char {
			return false
		}
	}
	return true
}

func isDiagonalBottomLeft(word []rune, x int, y int, positionsMap *map[int]map[int]rune, diff int) bool {
	for i, char := range word {
		row, ok := (*positionsMap)[x+i+diff]
		if !ok {
			return false
		}

		if ch, ok := row[y-i-diff]; !ok || ch != char {
			return false
		}
	}
	return true
}

func isDiagonalTopLeft(word []rune, x int, y int, positionsMap *map[int]map[int]rune, diff int) bool {
	for i, char := range word {
		row, ok := (*positionsMap)[x-i-diff]
		if !ok {
			return false
		}

		if ch, ok := row[y-i-diff]; !ok || ch != char {
			return false
		}
	}
	return true
}
