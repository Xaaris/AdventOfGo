package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
)

func solve1(inputPath string) int {
	measurements := IOUtils.ReadInputInts(inputPath)
	var depthIncreaseCounter int
	var previousDepth int
	for i, depth := range measurements {
		if i > 0 && depth > previousDepth {
			depthIncreaseCounter++
		}
		previousDepth = depth
	}

	return depthIncreaseCounter
}

func solve2(inputPath string) int {
	measurements := IOUtils.ReadInputInts(inputPath)
	var depthIncreaseCounter int
	for i := 0; i < len(measurements)-3; i++ {
		depthWindow := measurements[i+1] + measurements[i+2] + measurements[i+3]
		previousDepthWindow := measurements[i] + measurements[i+1] + measurements[i+2]
		if depthWindow > previousDepthWindow {
			depthIncreaseCounter++
		}
	}

	return depthIncreaseCounter
}

func main() {
	day := "day01"
	exampleResult := solve1(day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 7, exampleResult)
	result1 := solve1(day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2(day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 5, exampleResult)
	result2 := solve2(day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
