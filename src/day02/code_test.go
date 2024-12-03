package day02

import (
	"testing"
)

var test_input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	expected := 2
	data := parseInput(test_input)

	result := part1(data)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

func TestRemoveElement(t *testing.T) {
	expected := []int{2, 3, 4, 5}
	data := []int{1, 2, 3, 4, 5}

	result := removeElement(data, 0)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}

}

func TestPart2(t *testing.T) {
	expected := 4

	data := parseInput(test_input)

	result := part2(data)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}
