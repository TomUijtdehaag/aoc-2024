package utils

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput(d int) string {
	data, err := os.ReadFile(fmt.Sprintf("input/%02d.txt", d))
	check(err)

	return string(data)
}

func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}

func Split(input string, sep string) []string {
	return strings.Split(input, sep)
}

func ToInt(input string) int {
	var result int
	_, err := fmt.Sscanf(input, "%d", &result)
	check(err)

	return result
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
