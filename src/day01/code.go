package day01

import (
	"aoc/utils"
	"fmt"
	"sort"
)

func parseInput(input string) ([]int, []int) {
	lines := utils.SplitLines(input)
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		parts := utils.Split(line, "   ")
		left = append(left, utils.ToInt(parts[0]))
		right = append(right, utils.ToInt(parts[1]))
	}

	return left, right
}

func part1(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	diff := 0

	for i := 0; i < len(left); i++ {
		diff += utils.AbsInt(right[i] - left[i])
	}

	return diff
}

func part2(left, right []int) int {
	total := 0
	for l := range left {
		for r := range right {
			if left[l] == right[r] {
				total += left[l]
			}
		}
	}
	return total
}

func Solve() {
	input := utils.ReadInput(1)
	left, right := parseInput(input)

	fmt.Printf("Day 1, Part 1: %v\n", part1(left, right))
	fmt.Printf("Day 1, Part 2: %v\n", part2(left, right))

}
