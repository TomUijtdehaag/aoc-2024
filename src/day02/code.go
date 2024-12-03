package day02

import (
	"aoc/utils"
	"fmt"
)

func parseInput(input string) [][]int {
	lines := utils.SplitLines(input)
	result := make([][]int, len(lines))

	for i, line := range lines {
		for _, num := range utils.Split(line, " ") {
			result[i] = append(result[i], utils.ToInt(num))
		}
	}

	return result
}

func diffs(row []int) []int {
	result := make([]int, len(row)-1)
	for i := 0; i < len(row)-1; i++ {
		result[i] = row[i] - row[i+1]
	}
	return result
}

func isSafe(row []int) bool {
	anyIncreasing := false
	anyDecreasing := false
	outOfRange := false

	d := diffs(row)
	for _, diff := range d {
		if utils.AbsInt(diff) < 1 || utils.AbsInt(diff) > 3 {
			outOfRange = true
			break
		}
		if diff > 0 {
			anyIncreasing = true
		} else if diff < 0 {
			anyDecreasing = true
		}
	}

	if outOfRange || (anyIncreasing && anyDecreasing) {
		return false
	} else {
		return true
	}
}

func part1(input [][]int) int {
	result := 0
	for _, row := range input {
		if isSafe(row) {
			result++
		}

	}
	return result
}

func removeElement(slice []int, index int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	if index < 0 || index >= len(sliceCopy) {
		return sliceCopy
	}
	return append(sliceCopy[:index], sliceCopy[index+1:]...)
}

func part2(input [][]int) int {
	result := 0

	for _, row := range input {
		if isSafe(row) {
			result++
			continue
		}

		for i := 0; i < len(row); i++ {
			adjustedRow := removeElement(row, i)
			if isSafe(adjustedRow) {
				result++
				break
			}

		}
	}
	return result
}

func Solve() {
	input := utils.ReadInput(2)
	data := parseInput(input)

	fmt.Printf("Day 2, Part 1: %v\n", part1(data))
	fmt.Printf("Day 2, Part 2: %v\n", part2(data))
}
