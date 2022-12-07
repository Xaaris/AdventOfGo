package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
)

func solve1(inputPath string) int {
	input := IOUtils.ReadInputStrings(inputPath)[0]
	for i := 0; i < len(input); i++ {
		start := max(0, i-3)
		end := i
		set := map[rune]bool{}
		for _, r := range input[start : end+1] {
			set[r] = true
		}
		if len(set) == 4 {
			return i + 1
		}
	}

	return 0
}

func solve2(inputPath string) int {
	input := IOUtils.ReadInputStrings(inputPath)[0]
	for i := 0; i < len(input); i++ {
		start := max(0, i-13)
		end := i
		set := map[rune]bool{}
		for _, r := range input[start : end+1] {
			set[r] = true
		}
		if len(set) == 14 {
			return i + 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	day := "day06"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 11, exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	//exampleResult = solve2("2022/" + day + "/example.txt")
	//fmt.Printf("expected: %d, got %d\n", 0, exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
