package day04

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Direction int

const (
	up Direction = iota
	upRight
	right
	downRight
	down
	downLeft
	left
	upLeft
)

func parseInput(input string) [][]string {
	output := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		output = append(output, strings.Split(line, ""))
	}
	return output
}

func isInBounds(row, col int, grid [][]string) bool {
	return (row >= 0 && row < len(grid)) && (col >= 0 && col < len(grid[0]))
}

func getNext(row, col int, grid [][]string, direction Direction) (int, int) {
	switch direction {
	case up:
		return row - 1, col
	case upRight:
		return row - 1, col + 1
	case right:
		return row, col + 1
	case downRight:
		return row + 1, col + 1
	case down:
		return row + 1, col
	case downLeft:
		return row + 1, col - 1
	case left:
		return row, col - 1
	case upLeft:
		return row - 1, col - 1
	}
	return row, col
}

func checkWord(grid [][]string, srow, scol int, word string, direction Direction, wordIndex int) bool {
	row, col := srow, scol
	if string(word[wordIndex]) != grid[row][col] {
		return false
	}

	if wordIndex == len(word)-1 {
		return true
	}

	row, col = getNext(row, col, grid, direction)
	if !isInBounds(row, col, grid) {
		return false
	}

	return checkWord(grid, row, col, word, direction, wordIndex+1)
}

func part1(input string) int {
	grid := parseInput(input)
	total := 0
	for i, row := range grid {
		for j, letter := range row {
			if letter == "X" {
				for direction := up; direction <= upLeft; direction++ {
					if checkWord(grid, i, j, "XMAS", direction, 0) {
						total++
					}
				}
			}
		}
	}

	return total
}

func part2(input string) int {
	grid := parseInput(input)
	total := 0
	for i, row := range grid {
		for j, letter := range row {
			if letter == "A" {
				count := 0
				for direction := upRight; direction <= upLeft; direction += 2 {
					srow, scol := getNext(i, j, grid, direction)
					if !isInBounds(srow, scol, grid) {
						continue
					}
					if checkWord(grid, srow, scol, "MAS", (direction+4)%8, 0) {
						count++
					}
				}
				if count == 2 {
					total++
				}
			}
		}
	}
	return total
}

func Solve() {
	input := utils.ReadInput(4)

	fmt.Printf("Day 4, Part 1: %v\n", part1(input))
	fmt.Printf("Day 4, Part 2: %v\n", part2(input))
}
