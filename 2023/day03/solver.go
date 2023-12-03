package main

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	"AdventOfGo/IOUtils"
)

func solve1(inputPath string) int {
	enginePartSum := 0
	regex := regexp.MustCompile(`\d+`)
	lines := IOUtils.ReadInputStrings(inputPath)
	for lineNumber, line := range lines {
		matchIndexes := regex.FindAllStringIndex(line, -1)
		matches := regex.FindAllString(line, -1)
		for i, matchIndex := range matchIndexes {
			if hasAdjacentSymbol(inputPath, lineNumber, matchIndex[0], matchIndex[1]) {
				enginePartNumber, _ := strconv.Atoi(matches[i])
				enginePartSum += enginePartNumber
				fmt.Printf("Found engine part %d at line %d, column %d to %d\n", enginePartNumber, lineNumber, matchIndex[0], matchIndex[1])
			}
		}
	}
	return enginePartSum
}

func hasAdjacentSymbol(inputPath string, lineNumber, start, end int) bool {
	runes := IOUtils.ReadInputRunes(inputPath)
	minLineNumber := max(lineNumber-1, 0)
	maxLineNumber := min(lineNumber+1, len(runes)-1)
	minStart := max(start-1, 0)
	maxEnd := min(end, len(runes[lineNumber])-1)
	//for line := minLineNumber; line <= maxLineNumber; line++ {
	//	for i := minStart; i <= maxEnd; i++ {
	//		r := runes[line][i]
	//		fmt.Print(string(r))
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()
	for line := minLineNumber; line <= maxLineNumber; line++ {
		for i := minStart; i <= maxEnd; i++ {
			r := runes[line][i]
			if !unicode.IsDigit(r) && r != '.' {
				fmt.Printf("%c is symbol\n", r)
				return true
			}
		}
	}
	return false
}

func solve2(inputPath string) int {
	gearRatios := 0
	gears := make(map[int]map[int][]int)
	regex := regexp.MustCompile(`\d+`)
	lines := IOUtils.ReadInputStrings(inputPath)
	for lineNumber, line := range lines {
		matchIndexes := regex.FindAllStringIndex(line, -1)
		matches := regex.FindAllString(line, -1)
		for i, matchIndex := range matchIndexes {
			if line, col, hasGear := hasAdjacentGear(inputPath, lineNumber, matchIndex[0], matchIndex[1]); hasGear {
				gearRatioPart, _ := strconv.Atoi(matches[i])
				if gears[line] == nil {
					gears[line] = make(map[int][]int)
				}
				if gears[line][col] == nil {
					gears[line][col] = make([]int, 0)
				}
				gears[line][col] = append(gears[line][col], gearRatioPart)
				fmt.Printf("Found gear part %d with gear at line %d, column %d\n", gearRatioPart, line, col)
			}
		}
	}
	for _, line := range gears {
		for _, gearRatioParts := range line {
			if len(gearRatioParts) == 2 {
				gearRatios += gearRatioParts[0] * gearRatioParts[1]
			}
		}
	}
	return gearRatios
}

func hasAdjacentGear(inputPath string, lineNumber, start, end int) (int, int, bool) {
	runes := IOUtils.ReadInputRunes(inputPath)
	minLineNumber := max(lineNumber-1, 0)
	maxLineNumber := min(lineNumber+1, len(runes)-1)
	minStart := max(start-1, 0)
	maxEnd := min(end, len(runes[lineNumber])-1)
	for line := minLineNumber; line <= maxLineNumber; line++ {
		for i := minStart; i <= maxEnd; i++ {
			r := runes[line][i]
			if r == '*' {
				return line, i, true
			}
		}
	}
	return 0, 0, false
}

func main() {
	day := "day03"
	exampleResult := solve1("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 4361, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 467835, exampleResult)
	result2 := solve2("2023/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
