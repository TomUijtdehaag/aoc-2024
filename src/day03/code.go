package day03

import (
	"aoc/utils"
	"fmt"
	"regexp"
)

func parseInput(input string) [][]int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	output := [][]int{}

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		left, right := match[1], match[2]

		output = append(output, []int{utils.ToInt(left), utils.ToInt(right)})
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
	pairs := parseInput(input)
	return multiplyPairs(pairs)
}

func cleanInput(input string) string {
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	return re.ReplaceAllString(input+"do()", "")
}

func part2(input string) int {
	input = cleanInput(input)
	pairs := parseInput(input)
	return multiplyPairs(pairs)
}

func Solve() {
	input := utils.ReadInput(3)

	fmt.Printf("Day 2, Part 1: %v\n", part1(input))
	fmt.Printf("Day 2, Part 2: %v\n", part2(input))
}
