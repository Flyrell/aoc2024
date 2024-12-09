package day4

import (
	"github.com/flyrell/aoc2024/internal/utils"
)

func RunExtra() int {
	sum := 0
	word := []rune{'M', 'A', 'S'}
	aPositions := make([][]int, 0, 20)
	positionsMap := make(map[int]map[int]rune)
	utils.ReadInputFileByLine("/app/internal/day4/input", func(rowNo int, text string) {
		for colNo, char := range text {
			if char != word[0] && char != word[1] && char != word[2] {
				continue
			}

			if _, ok := positionsMap[rowNo]; !ok {
				positionsMap[rowNo] = make(map[int]rune)
			}
			positionsMap[rowNo][colNo] = char

			if char == word[1] {
				aPositions = append(aPositions, []int{rowNo, colNo})
				continue
			}
		}
	})

	for _, coords := range aPositions {
		x := coords[0]
		y := coords[1]

		isRightDiagonal := isDiagonalBottomRight(word, x-1, y-1, &positionsMap, 0) || isDiagonalTopLeft(word, x+1, y+1, &positionsMap, 0)
		isLeftDiagonal := isDiagonalBottomLeft(word, x-1, y+1, &positionsMap, 0) || isDiagonalTopRight(word, x+1, y-1, &positionsMap, 0)

		if isRightDiagonal && isLeftDiagonal {
			sum++
		}
	}

	defer func() {
		println("Result of Day 4 Extra is:", sum)
	}()

	return sum
}
