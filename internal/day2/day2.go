package day2

import (
	"github.com/flyrell/aoc2024/internal/utils"
	"math"
	"strconv"
	"strings"
)

func Run() int {
	sum := 0
	utils.ReadInputFileByLine("/app/internal/day2/input", func(_ int, row string) {
		levelsAsStrings := strings.Split(row, " ")
		levels := make([]int, 0, 10)
		for _, level := range levelsAsStrings {
			num, err := strconv.Atoi(level)
			if err == nil {
				levels = append(levels, num)
			}
		}

		direction := levels[0] > levels[1]
		lastNumber := levels[0]
		for _, number := range levels[1:] {
			isDecreasing := direction && lastNumber > number
			isIncreasing := !direction && lastNumber < number
			if !isDecreasing && !isIncreasing {
				return
			}

			abs := math.Abs(float64(lastNumber - number))
			if abs < 1 || abs > 3 {
				return
			}

			lastNumber = number
		}

		sum += 1
	})

	defer func() {
		println("Result of Day 2 is:", sum)
	}()

	return sum
}
