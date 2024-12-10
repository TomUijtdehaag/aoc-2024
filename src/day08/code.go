package day08

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Location struct {
	x int
	y int
}

type LocationPair struct {
	locationA Location
	locationB Location
}

type Map struct {
	grid      []string
	locations map[string][]Location
	antiNodes [][]bool
}

func NewMap(input string) *Map {
	grid := strings.Split(input, "\n")
	locations := make(map[string][]Location)
	antiNodes := make([][]bool, len(grid))

	for y, row := range grid {
		antiNodes[y] = make([]bool, len(row))
		for x, char := range row {
			if char != '.' {
				locations[string(char)] = append(locations[string(char)], Location{x, y})
			}
		}
	}

	return &Map{grid: grid, locations: locations, antiNodes: antiNodes}
}

func (m *Map) GetPairs(frequency string) []LocationPair {
	locations := m.locations[frequency]
	pairs := []LocationPair{}

	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			pairs = append(pairs, LocationPair{locations[i], locations[j]})
		}
	}

	return pairs
}

func (m *Map) FindAntiNodes(locationPair LocationPair) []Location {
	A, B := locationPair.locationA, locationPair.locationB
	xDiff := B.x - A.x
	yDiff := B.y - A.y

	node1 := Location{B.x + xDiff, B.y + yDiff}
	node2 := Location{A.x - xDiff, A.y - yDiff}

	return []Location{node1, node2}
}

func (m *Map) FindAntiNodesResonant(locationPair LocationPair) []Location {
	A, B := locationPair.locationA, locationPair.locationB
	xDiff := B.x - A.x
	yDiff := B.y - A.y

	nodes := []Location{}

	steps := 0
	for {
		node := Location{B.x + xDiff*steps, B.y + yDiff*steps}
		if m.IsInGrid(node) {
			nodes = append(nodes, node)
		} else {
			break
		}
		steps++
	}

	steps = 0
	for {
		node := Location{A.x - xDiff*steps, A.y - yDiff*steps}
		if m.IsInGrid(node) {
			nodes = append(nodes, node)
		} else {
			break
		}
		steps++
	}

	return nodes
}

func (m *Map) CountAntiNodes() int {
	count := 0

	for _, row := range m.antiNodes {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}

	return count
}

func (m *Map) IsInGrid(location Location) bool {
	return location.x >= 0 && location.x < len(m.grid[0]) && location.y >= 0 && location.y < len(m.grid)
}

func (m *Map) PrintGrid() {
	for _, row := range m.grid {
		fmt.Println(row)
	}
}

func (m *Map) PrintAntiNodes() {
	for _, row := range m.antiNodes {
		for _, cell := range row {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part1(input string) int {
	m := NewMap(input)

	for frequency, _ := range m.locations {
		for _, locationPair := range m.GetPairs(frequency) {
			antiNodes := m.FindAntiNodes(locationPair)
			for _, antiNode := range antiNodes {
				if m.IsInGrid(antiNode) {
					m.antiNodes[antiNode.y][antiNode.x] = true
				}
			}
		}
	}

	m.PrintGrid()
	fmt.Println()
	m.PrintAntiNodes()

	return m.CountAntiNodes()
}

func part2(input string) int {
	m := NewMap(input)

	for frequency, _ := range m.locations {
		for _, locationPair := range m.GetPairs(frequency) {
			antiNodes := m.FindAntiNodesResonant(locationPair)
			for _, antiNode := range antiNodes {
				m.antiNodes[antiNode.y][antiNode.x] = true
			}
		}
	}

	m.PrintGrid()
	fmt.Println()
	m.PrintAntiNodes()

	return m.CountAntiNodes()
}

func Solve() {
	input := utils.ReadInput(8)

	fmt.Printf("Day 7, Part 1: %v\n", part1(input))
	fmt.Printf("Day 7, Part 2: %v\n", part2(input))
}
