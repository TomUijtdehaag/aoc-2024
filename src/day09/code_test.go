package day09

import "testing"

const testInput = `2333133121414131402`

func TestPart1(t *testing.T) {
	expected := 1928
	result := part1(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 2858
	result := part2(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
