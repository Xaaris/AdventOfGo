package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func solve1(inputPath string) int {
	commands := IOUtils.ReadInputStrings(inputPath)
	var depth int
	var horizontalPosition int

	for _, command := range commands {
		parts := strings.Split(command, " ")
		direction := parts[0]
		value, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			horizontalPosition += value
		case "up":
			depth -= value
		case "down":
			depth += value
		default:
			log.Fatal("Unknown direction: " + direction)
		}
	}
	return horizontalPosition * depth
}

func solve2(inputPath string) int {
	commands := IOUtils.ReadInputStrings(inputPath)
	var aim int
	var depth int
	var horizontalPosition int

	for _, command := range commands {
		parts := strings.Split(command, " ")
		direction := parts[0]
		value, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			horizontalPosition += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		default:
			log.Fatal("Unknown direction: " + direction)
		}
	}
	return horizontalPosition * depth
}

func main() {
	day := "day02"
	exampleResult := solve1(day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 150, exampleResult)
	result1 := solve1(day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2(day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 900, exampleResult)
	result2 := solve2(day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
