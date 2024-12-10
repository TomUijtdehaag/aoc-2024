package day08

import "testing"

const testInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestPart1(t *testing.T) {
	expected := 14
	result := part1(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 34
	result := part2(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
