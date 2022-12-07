package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"strconv"
	"strings"
)

func solve1(inputPath string) int {
	fullyContainedCounter := 0
	input := IOUtils.ReadInputStrings(inputPath)
	for _, pair := range input {
		a, b := toSections(pair)
		if (a.lower >= b.lower && a.upper <= b.upper) || (b.lower >= a.lower && b.upper <= a.upper) {
			fullyContainedCounter++
		}
	}
	return fullyContainedCounter
}

func solve2(inputPath string) int {
	partiallyContainedCounter := 0
	input := IOUtils.ReadInputStrings(inputPath)
	for _, pair := range input {
		a, b := toSections(pair)
		if (a.lower >= b.lower && a.upper <= b.upper) ||
			(b.lower >= a.lower && b.upper <= a.upper) ||
			(a.lower <= b.lower && a.upper >= b.lower) ||
			(a.lower <= b.upper && a.upper >= b.upper) ||
			(b.lower <= a.lower && b.upper >= a.lower) ||
			(b.lower <= a.upper && b.upper >= a.upper) {
			partiallyContainedCounter++
		}
	}
	return partiallyContainedCounter
}

func toSections(pair string) (a, b section) {
	sections := strings.Split(pair, ",")
	a.lower, _ = strconv.Atoi(strings.Split(sections[0], "-")[0])
	a.upper, _ = strconv.Atoi(strings.Split(sections[0], "-")[1])
	b.lower, _ = strconv.Atoi(strings.Split(sections[1], "-")[0])
	b.upper, _ = strconv.Atoi(strings.Split(sections[1], "-")[1])
	return a, b
}

type section struct {
	lower int
	upper int
}

func main() {
	day := "day04"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 2, exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 4, exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
