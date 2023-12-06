package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfGo/IOUtils"
)

func solve1(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	times := stringsToInts(strings.Fields(lines[0][10:]))
	distances := stringsToInts(strings.Fields(lines[1][10:]))
	fmt.Printf("times: %v\n", times)
	fmt.Printf("distances: %v\n", distances)
	score := 1
	for raceNumber := 0; raceNumber < len(times); raceNumber++ {
		waysToWin := 0
		for buttonHoldTime := 1; buttonHoldTime < times[raceNumber]; buttonHoldTime++ {
			speed := buttonHoldTime
			timeRemaining := times[raceNumber] - buttonHoldTime
			distanceTraveled := speed * timeRemaining
			if distanceTraveled > distances[raceNumber] {
				waysToWin++
			}
		}
		score *= waysToWin
	}
	return score
}

func stringsToInts(strings []string) []int {
	ints := []int{}
	for _, str := range strings {
		atoi, _ := strconv.Atoi(str)
		ints = append(ints, atoi)
	}
	return ints
}

func solve2(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	time, _ := strconv.Atoi(strings.Join(strings.Fields(lines[0][10:]), ""))
	distance, _ := strconv.Atoi(strings.Join(strings.Fields(lines[1][10:]), ""))
	fmt.Printf("times: %v\n", time)
	fmt.Printf("distances: %v\n", distance)
	waysToWin := 0
	for buttonHoldTime := 1; buttonHoldTime < time; buttonHoldTime++ {
		speed := buttonHoldTime
		timeRemaining := time - buttonHoldTime
		distanceTraveled := speed * timeRemaining
		if distanceTraveled > distance {
			waysToWin++
		}
	}
	return waysToWin
}

func main() {
	day := "day06"
	exampleResult := solve1("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 288, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 71503, exampleResult)
	result2 := solve2("2023/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
