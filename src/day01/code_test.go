package day01

import (
	"testing"
)

var test_input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	expected := 11
	left, right := parseInput(test_input)

	result := part1(left, right)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

func TestPart2(t *testing.T) {
	expected := 31

	left, right := parseInput(test_input)

	result := part2(left, right)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}
