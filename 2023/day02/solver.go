package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"AdventOfGo/IOUtils"
)

func solve1(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)

	totalScore := 0
	for _, line := range lines {
		gameId, _ := strconv.Atoi(regexp.MustCompile(`(\d+):`).FindStringSubmatch(line)[1])
		roundsString := strings.Split(line, ":")[1]
		rounds := strings.Split(roundsString, ";")
		highestNumberPerColor := make(map[string]int)
		for _, round := range rounds {
			numberColorStrings := strings.Split(round, ",")
			for _, numberColorString := range numberColorStrings {
				number, _ := strconv.Atoi(regexp.MustCompile(`(\d+)`).FindStringSubmatch(numberColorString)[1])
				color := regexp.MustCompile(`([a-zA-Z]+)`).FindStringSubmatch(numberColorString)[1]
				if number > highestNumberPerColor[color] {
					highestNumberPerColor[color] = number
				}
			}

		}
		fmt.Println(gameId, highestNumberPerColor)
		if highestNumberPerColor["red"] <= 12 &&
			highestNumberPerColor["green"] <= 13 &&
			highestNumberPerColor["blue"] <= 14 {
			totalScore += gameId
		}
	}
	return totalScore
}

func solve2(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)

	totalScore := 0
	for _, line := range lines {
		gameId, _ := strconv.Atoi(regexp.MustCompile(`(\d+):`).FindStringSubmatch(line)[1])
		roundsString := strings.Split(line, ":")[1]
		rounds := strings.Split(roundsString, ";")
		highestNumberPerColor := make(map[string]int)
		for _, round := range rounds {
			numberColorStrings := strings.Split(round, ",")
			for _, numberColorString := range numberColorStrings {
				number, _ := strconv.Atoi(regexp.MustCompile(`(\d+)`).FindStringSubmatch(numberColorString)[1])
				color := regexp.MustCompile(`([a-zA-Z]+)`).FindStringSubmatch(numberColorString)[1]
				if number > highestNumberPerColor[color] {
					highestNumberPerColor[color] = number
				}
			}

		}
		fmt.Println(gameId, highestNumberPerColor)
		power := highestNumberPerColor["red"] * highestNumberPerColor["green"] * highestNumberPerColor["blue"]
		totalScore += power
	}
	return totalScore
}

func main() {
	day := "day02"
	exampleResult := solve1("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 8, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 2286, exampleResult)
	result2 := solve2("2023/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
