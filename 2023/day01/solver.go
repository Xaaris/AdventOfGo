package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"AdventOfGo/IOUtils"
)

func solve1(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	calibrationValue := 0
	for _, line := range lines {
		var firstDigit string
		var lastDigit string
		runes := []rune(line)
		for _, c := range runes {
			if unicode.IsDigit(c) {
				firstDigit = string(c)
				break
			}
		}
		slices.Reverse(runes)
		for _, c := range runes {
			if unicode.IsDigit(c) {
				lastDigit = string(c)
				break
			}
		}

		combinedNumbersAsString := fmt.Sprintf("%s%s", firstDigit, lastDigit)
		combinedNumbers, _ := strconv.Atoi(combinedNumbersAsString)
		calibrationValue += combinedNumbers
	}
	return calibrationValue
}

func solve2(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)

	writtenNumbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	calibrationValue := 0
	for _, line := range lines {
		var firstDigit string
		firstDigitFound := false
		var lastDigit string
		runes := []rune(line)

		for i, c := range runes {
			if unicode.IsDigit(c) {
				if !firstDigitFound {
					firstDigit = string(c)
					firstDigitFound = true
				}
				lastDigit = string(c)
			}
			for letters, num := range writtenNumbers {
				if strings.HasPrefix(line[i:], letters) {
					if !firstDigitFound {
						firstDigit = num
						firstDigitFound = true
					}
					lastDigit = num
				}
			}
		}
		combinedNumbersAsString := fmt.Sprintf("%s%s", firstDigit, lastDigit)
		combinedNumbers, _ := strconv.Atoi(combinedNumbersAsString)
		calibrationValue += combinedNumbers
	}
	return calibrationValue
}

func main() {
	day := "day01"
	exampleResult := solve1("2023/" + day + "/example1.txt")
	fmt.Printf("expected: %d, got %d\n", 142, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)
	//
	exampleResult = solve2("2023/" + day + "/example2.txt")
	fmt.Printf("expected: %d, got %d\n", 281, exampleResult)
	result2 := solve2("2023/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
