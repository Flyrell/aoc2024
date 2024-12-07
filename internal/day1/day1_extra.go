package day1

import (
	"github.com/flyrell/aoc2024/internal/utils"
	"strconv"
	"strings"
)

func RunExtra() int {
	var left = make([]int, 1000)
	var right = make(map[int]int)
	utils.ReadInputFileByLine("/app/internal/day1/input", func(i int, row string) {
		numbers := strings.Split(row, "   ")

		leftN, err := strconv.Atoi(numbers[0])
		if err == nil {
			left[i] = leftN
		}

		rightN, err := strconv.Atoi(numbers[1])
		if err != nil {
			return
		}

		if _, ok := right[rightN]; !ok {
			right[rightN] = 1
		} else {
			right[rightN] += 1
		}
	})

	sum := 0
	for _, num := range left {
		if _, ok := right[num]; !ok {
			continue
		}
		sum += num * right[num]
	}

	defer func() {
		println("Result of Day 1 Extra is:", sum)
	}()

	return sum
}
