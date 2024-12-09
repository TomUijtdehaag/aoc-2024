package day06

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Map struct {
	obstacles  [][]bool
	visited    [][]bool
	directions [][]Direction
	position   Point
	direction  Direction
}

func (m *Map) Parse(input string) {
	lines := strings.Split(input, "\n")
	m.obstacles = make([][]bool, len(lines))
	m.visited = make([][]bool, len(lines))
	m.directions = make([][]Direction, len(lines))
	for y, line := range lines {
		m.obstacles[y] = make([]bool, len(line))
		m.visited[y] = make([]bool, len(line))
		m.directions[y] = make([]Direction, len(line))
		for x, char := range line {
			switch char {
			case '#':
				m.obstacles[y][x] = true
			case '^':
				m.position = Point{x, y}
				m.direction = North
				m.visited[y][x] = true
				m.directions[y][x] = North
			}
		}
	}
}

func (m *Map) Peek(direction Direction) Point {
	switch direction {
	case North:
		return Point{m.position.x, m.position.y - 1}
	case East:
		return Point{m.position.x + 1, m.position.y}
	case South:
		return Point{m.position.x, m.position.y + 1}
	case West:
		return Point{m.position.x - 1, m.position.y}
	}
	return Point{-1, -1}
}

func (m *Map) IsObstacle(point Point) bool {
	return m.obstacles[point.y][point.x]
}

func (m *Map) Turn() Direction {
	return (m.direction + 1) % 4
}

func (m *Map) outOfBounds(point Point) bool {
	return point.x < 0 || point.x >= len(m.obstacles[0]) || point.y < 0 || point.y >= len(m.obstacles)
}

func (m *Map) Step() bool {
	nextPosition := m.Peek(m.direction)
	if m.outOfBounds(nextPosition) {
		return true
	}
	for m.IsObstacle(nextPosition) {
		m.direction = m.Turn()
		nextPosition = m.Peek(m.direction)
	}
	m.position = nextPosition
	m.visited[m.position.y][m.position.x] = true
	m.directions[m.position.y][m.position.x] = m.direction
	return false
}

func (m *Map) Sum() int {
	sum := 0
	for _, row := range m.visited {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}
	return sum
}

func (m *Map) SetObstacle(point Point) {
	m.obstacles[point.y][point.x] = true
}

func NewMap(input string) *Map {
	m := &Map{}
	m.Parse(input)
	return m
}

func part1(input string) int {
	m := NewMap(input)
	for !m.Step() {
	}

	return m.Sum()
}

func part2(input string) int {
	solutions := 0
	startMap := NewMap(input)

	for y, row := range startMap.obstacles {
		for x, obstacle := range row {
			if obstacle || startMap.position.x == x && startMap.position.y == y {
				continue
			}

			m := NewMap(input)
			m.SetObstacle(Point{x, y})
			for !m.Step() {
				nextPosition := m.Peek(m.direction)
				if m.outOfBounds(nextPosition) {
					break
				}
				if m.visited[nextPosition.y][nextPosition.x] && m.directions[nextPosition.y][nextPosition.x] == m.direction {
					solutions++
					break
				}
			}
		}
	}

	return solutions
}

func Solve() {
	input := utils.ReadInput(6)

	fmt.Printf("Day 6, Part 1: %v\n", part1(input))
	fmt.Printf("Day 6, Part 2: %v\n", part2(input))
}
