package day06

import "testing"

const testInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	expected := 41
	result := part1(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 6
	result := part2(testInput)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
