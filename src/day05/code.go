package day05

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Rule struct {
	a int
	b int
}

func NewRule(input string) Rule {
	parts := strings.Split(input, "|")
	return Rule{a: utils.ToInt(parts[0]), b: utils.ToInt(parts[1])}
}

type Rules []Rule

func NewRules(input []string) Rules {
	rules := make(Rules, len(input))
	for i, line := range input {
		rules[i] = NewRule(line)
	}
	return rules
}

type Update struct {
	pages   []int
	pagemap map[int]int
}

func NewUpdate(input string) Update {
	pagesInput := strings.Split(input, ",")
	pages := make([]int, len(pagesInput))
	for i, page := range pagesInput {
		pages[i] = utils.ToInt(page)
	}
	pagemap := make(map[int]int)
	for i, page := range pages {
		pagemap[page] = i
	}
	return Update{pages: pages, pagemap: pagemap}
}

func (u Update) Check(rule Rule) bool {
	a, aok := u.pagemap[rule.a]
	b, bok := u.pagemap[rule.b]
	if !(aok && bok) {
		return true
	}
	if a > b {
		return false
	}
	return true
}

func (u Update) CheckAll(rules Rules) (bool, Rule) {
	for _, rule := range rules {
		if !u.Check(rule) {
			return false, rule
		}
	}
	return true, Rule{}
}

func (u Update) Swap(rule Rule) {
	a, aok := u.pagemap[rule.a]
	b, bok := u.pagemap[rule.b]
	if !(aok && bok) {
		return
	}
	u.pages[a], u.pages[b] = u.pages[b], u.pages[a]
	u.pagemap[rule.a], u.pagemap[rule.b] = b, a
}

type Updates []Update

func NewUpdates(input []string) Updates {
	updates := make(Updates, len(input))
	for i, line := range input {
		updates[i] = NewUpdate(line)
	}
	return updates
}

func parseInput(input string) (Rules, Updates) {
	inputSplit := strings.Split(input, "\n\n")
	rulesInput := strings.Split(inputSplit[0], "\n")
	updatesInput := strings.Split(inputSplit[1], "\n")

	rules := NewRules(rulesInput)
	updates := NewUpdates(updatesInput)

	return rules, updates
}

func part1(input string) int {
	rules, updates := parseInput(input)

	total := 0
	for _, update := range updates {
		if ok, _ := update.CheckAll(rules); ok {
			total += update.pages[len(update.pages)/2]
		}
	}

	return total
}

func part2(input string) int {
	rules, updates := parseInput(input)

	total := 0
	for _, update := range updates {
		ok, _ := update.CheckAll(rules)
		if !ok {
			for {
				ok, rule := update.CheckAll(rules)
				if ok {
					break
				}
				update.Swap(rule)
			}
			total += update.pages[len(update.pages)/2]
		}
	}

	return total
}

func Solve() {
	input := utils.ReadInput(5)

	fmt.Printf("Day 5, Part 1: %v\n", part1(input))
	fmt.Printf("Day 5, Part 2: %v\n", part2(input))
}
