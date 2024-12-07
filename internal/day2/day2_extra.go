package day2

import (
	"github.com/flyrell/aoc2024/internal/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func check(levels []int) bool {
	direction := levels[0] > levels[1]
	lastNumber := levels[0]
	for _, number := range levels[1:] {
		isDecreasing := direction && lastNumber > number
		isIncreasing := !direction && lastNumber < number
		if !isDecreasing && !isIncreasing {
			return false
		}

		abs := math.Abs(float64(lastNumber - number))
		if abs < 1 || abs > 3 {
			return false
		}

		lastNumber = number
	}

	return true
}

func RunExtra() int {
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

		if check(levels) {
			sum += 1
			return
		}

		for i := 0; i < len(levels); i++ {
			if check(slices.Concat(nil, levels[:i], levels[i+1:])) {
				sum += 1
				break
			}
		}
	})

	defer func() {
		println("Result of Day 2 Extra is:", sum)
	}()

	return sum
}
