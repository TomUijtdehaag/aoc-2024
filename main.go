package main

import (
	"aoc/src/day01"
	"aoc/src/day02"
	"aoc/src/day03"
	"aoc/src/day04"
	"aoc/src/day05"
	"aoc/src/day06"
	"aoc/src/day07"
	"aoc/src/day08"
	"aoc/src/day09"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a day number as argument")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day number:", os.Args[1])
		os.Exit(1)
	}

	switch day {
	case 1:
		day01.Solve()
	case 2:
		day02.Solve()
	case 3:
		day03.Solve()
	case 4:
		day04.Solve()
	case 5:
		day05.Solve()
	case 6:
		day06.Solve()
	case 7:
		day07.Solve()
	case 8:
		day08.Solve()
	case 9:
		day09.Solve(day)
	default:
		fmt.Printf("Day %d not implemented\n", day)
		os.Exit(1)
	}
}
