package day1

import (
	"github.com/flyrell/aoc2024/internal/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Run() float64 {
	var left = make([]int, 1000)
	var right = make([]int, 1000)
	utils.ReadInputFileByLine("/app/internal/day1/input", func(i int, row string) {
		numbers := strings.Split(row, "   ")

		leftN, err := strconv.Atoi(numbers[0])
		if err == nil {
			left[i] = leftN
		}

		rightN, err := strconv.Atoi(numbers[1])
		if err == nil {
			right[i] = rightN
		}
	})

	sum := 0.0
	sort.Ints(left)
	sort.Ints(right)
	for i := range left {
		sum += math.Abs(float64(left[i] - right[i]))
	}

	defer func() {
		println("Result of Day 1 is:", int(sum))
	}()
	return sum
}
