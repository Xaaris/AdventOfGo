package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"strings"
)

func solve1(inputPath string) int {
	totalScore := 0
	rounds := IOUtils.ReadInputStrings(inputPath)
	for _, round := range rounds {
		hands := strings.Split(round, " ")
		you := toNumber(hands[0])
		me := toNumber(hands[1])
		winScore := determineWinScore(you, me)
		totalScore += winScore + me
	}
	return totalScore
}

func solve2(inputPath string) int {
	totalScore := 0
	rounds := IOUtils.ReadInputStrings(inputPath)
	for _, round := range rounds {
		hands := strings.Split(round, " ")
		you := toNumber(hands[0])
		outcome := toNumber(hands[1])
		me := calculateOwnHand(you, outcome)
		winScore := determineWinScore(you, me)
		totalScore += winScore + me
	}
	return totalScore
}

func calculateOwnHand(you int, outcome int) int {
	switch outcome {
	case 1:
		me := you - 1
		if me == 0 {
			me = 3
		}
		return me
	case 2:
		return you
	case 3:
		me := you + 1
		if me == 4 {
			me = 1
		}
		return me
	default:
		panic("unexpected outcome")
	}
	return 0
}

func determineWinScore(you, me int) int {
	if you == me {
		return 3
	}
	if you+1 == 4 {
		you = 0
	}
	if you+1 == me {
		return 6
	}
	return 0
}

func toNumber(char string) int {
	switch char {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	default:
		panic("unexpected hand")
	}

	return 0
}

func main() {
	day := "day02"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 15, exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 12, exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
