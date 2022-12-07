package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"strconv"
)

func solve1(inputPath string) int {
	var gamma, epsilon int64
	runeLines := IOUtils.ReadInputRunes(inputPath)
	floats := convertToFloat(runeLines)
	averageBits := make([]float32, len(floats[0]))
	for _, line := range floats {
		for i, bit := range line {
			averageBits[i] += bit / float32(len(floats))
		}
	}
	gammaString, epsilonString := "", ""
	for _, bit := range averageBits {
		if bit > 0.5 {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}
	gamma, _ = strconv.ParseInt(gammaString, 2, 0)
	epsilon, _ = strconv.ParseInt(epsilonString, 2, 0)

	return int(gamma * epsilon)
}

func solve2(inputPath string) int {
	return 0
}

func main() {
	day := "day03"
	exampleResult := solve1(day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 198, exampleResult)
	result1 := solve1(day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2(day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 900, exampleResult)
	result2 := solve2(day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}

func convertToFloat(runeLines [][]rune) [][]float32 {
	floats := make([][]float32, len(runeLines))
	for i := range floats {
		floats[i] = make([]float32, len(runeLines[0]))
	}

	for i, line := range runeLines {
		for j, r := range line {
			f, _ := strconv.Atoi(string(r))
			floats[i][j] = float32(f)
		}
	}
	return floats
}
