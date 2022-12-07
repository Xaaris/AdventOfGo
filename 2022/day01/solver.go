package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"strconv"
)

func solve1(inputPath string) int {
	calories := IOUtils.ReadInputStrings(inputPath)
	var aggregatedCalories []int
	aggregate := 0

	for _, entry := range calories {
		if entry == "" {
			aggregatedCalories = append(aggregatedCalories, aggregate)
			aggregate = 0
			continue
		}
		entryInt, _ := strconv.Atoi(entry)
		aggregate += entryInt
	}

	max := 0
	for _, entry := range aggregatedCalories {
		if entry > max {
			max = entry
		}
	}

	return max
}

func solve2(inputPath string) int {
	calories := IOUtils.ReadInputStrings(inputPath)
	var aggregatedCalories []int
	aggregate := 0

	for _, entry := range calories {
		if entry == "" {
			aggregatedCalories = append(aggregatedCalories, aggregate)
			aggregate = 0
			continue
		}
		entryInt, _ := strconv.Atoi(entry)
		aggregate += entryInt
	}
	aggregatedCalories = append(aggregatedCalories, aggregate)

	maxTotal, max := 0, 0
	for i := 0; i < 3; i++ {
		max, aggregatedCalories = extractMax(aggregatedCalories)
		maxTotal += max

	}

	return maxTotal
}

func extractMax(list []int) (int, []int) {
	max := 0
	index := 0
	for i, entry := range list {
		if entry > max {
			max = entry
			index = i
		}
	}
	list = remove(list, index)
	return max, list
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	day := "day01"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 24000, exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 45000, exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
