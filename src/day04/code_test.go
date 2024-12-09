package day04

import "testing"

const testInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {
	expected := 18
	result := part1(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 9
	result := part2(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
