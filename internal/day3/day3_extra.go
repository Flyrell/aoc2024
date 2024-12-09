package day3

import (
	"github.com/flyrell/aoc2024/internal/utils"
	"regexp"
	"strconv"
)

func RunExtra() int {
	sum := 0
	utils.ReadInputFileByLine("/app/internal/day3/input", func(_ int, input string) {
		enabled := true
		rg := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
		matches := rg.FindAllString(input, -1)

		numRg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
		for _, match := range matches {
			if match == "do()" {
				enabled = true
				continue
			}

			if !enabled {
				continue
			}

			if match == "don't()" {
				enabled = false
				continue
			}

			numbers := numRg.FindAllStringSubmatch(match, -1)
			if len(numbers) == 0 {
				continue
			}

			num1, err1 := strconv.Atoi(numbers[0][1])
			num2, err2 := strconv.Atoi(numbers[0][2])
			if err1 == nil && err2 == nil {
				sum += num1 * num2
			}
		}
	})

	defer func() {
		println("Result of Day 3 Extra is:", sum)
	}()

	return sum
}
