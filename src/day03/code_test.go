package day03

import "testing"

func TestPart1(t *testing.T) {

	testInput := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := part1(testInput)
	if result != 161 {
		t.Errorf("Expected 161, got %v", result)
	}
}

func TestPart2(t *testing.T) {
	testInput := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := part2(testInput)
	if result != 48 {
		t.Errorf("Expected 48, got %v", result)
	}
}
