package day07

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Operator int

const (
	Add Operator = iota
	Multiply
	Concatenate
)

type Equation struct {
	value   int
	numbers []int
}

func SolveStep(e Equation, op Operator) Equation {
	if len(e.numbers) < 2 {
		return e
	}

	left := e.numbers[0]
	right := e.numbers[1]
	remaining := e.numbers[2:]

	var result int
	switch op {
	case Add:
		result = left + right
	case Multiply:
		result = left * right
	case Concatenate:
		result = utils.ToInt(strings.Join([]string{strconv.Itoa(left), strconv.Itoa(right)}, ""))
	default:
		panic(fmt.Sprintf("Unknown operator: %v", op))
	}

	newNumbers := append([]int{result}, remaining...)
	return Equation{value: e.value, numbers: newNumbers}
}

func isSolved(e Equation) bool {
	return len(e.numbers) == 1 && e.value == e.numbers[0]
}

func isSolvable(e Equation, ops []Operator) bool {
	if len(e.numbers) == 1 {
		return e.numbers[0] == e.value
	}

	for _, op := range ops {
		newEquation := SolveStep(e, op)
		if isSolvable(newEquation, ops) {
			return true
		}
	}

	return false
}

func parseEquation(input string) Equation {
	parts := strings.Split(input, ": ")
	value := utils.ToInt(parts[0])
	numbers := strings.Split(parts[1], " ")

	equation := Equation{value, []int{}}
	for _, number := range numbers {
		equation.numbers = append(equation.numbers, utils.ToInt(number))
	}
	return equation
}

func parseInput(input string) []Equation {
	lines := strings.Split(input, "\n")
	equations := []Equation{}
	for _, line := range lines {
		equations = append(equations, parseEquation(line))
	}
	return equations
}

func part1(input string) int {
	equations := parseInput(input)
	total := 0

	ops := []Operator{Add, Multiply}

	for _, equation := range equations {
		if isSolvable(equation, ops) {
			total += equation.value
		}
	}

	return total
}

func part2(input string) int {
	equations := parseInput(input)
	total := 0

	ops := []Operator{Add, Multiply, Concatenate}

	for _, equation := range equations {
		if isSolvable(equation, ops) {
			total += equation.value
		}
	}

	return total
}

func Solve() {
	input := utils.ReadInput(7)

	fmt.Printf("Day 7, Part 1: %v\n", part1(input))
	fmt.Printf("Day 7, Part 2: %v\n", part2(input))
}
