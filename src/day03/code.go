package day03

import (
	"aoc/utils"
	"fmt"
	"regexp"
)

func parseInput(input string, do bool) [][]int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	output := [][]int{}
	mult := true

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		switch match[0] {
		case "do()":
			mult = true
		case "don't()":
			mult = false
		default:
			left, right := match[1], match[2]
			if do && !mult {
				continue
			} else {
				output = append(output, []int{utils.ToInt(left), utils.ToInt(right)})
			}
		}

	}
	return output
}

func multiplyPairs(pairs [][]int) int {
	total := 0
	for _, pair := range pairs {
		total += pair[0] * pair[1]
	}
	return total
}

func part1(input string) int {
	pairs := parseInput(input, false)
	return multiplyPairs(pairs)
}

func part2(input string) int {
	pairs := parseInput(input, true)
	return multiplyPairs(pairs)
}

func Solve() {
	input := utils.ReadInput(3)

	fmt.Printf("Day 2, Part 1: %v\n", part1(input))
	fmt.Printf("Day 2, Part 2: %v\n", part2(input))
}
