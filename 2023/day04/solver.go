package main

import (
	"fmt"
	"slices"
	"strings"

	"AdventOfGo/IOUtils"
)

func solve1(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	score := 0
	for _, line := range lines {
		cardScore := 0
		lineWithoutPrefix := strings.Replace(line, "Card 1: ", "", 1)
		pair := strings.Split(lineWithoutPrefix, "|")
		winningNumbers, myNumbers := strings.Fields(pair[0]), strings.Fields(pair[1])
		for _, myNumber := range myNumbers {
			if slices.Contains(winningNumbers, myNumber) {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}
		score += cardScore
	}
	return score
}

type Card struct {
	WinningNumbers []string
	MyNumbers      []string
	AmountOfCards  int
}

func solve2(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	cards := make([]Card, 0)
	for _, line := range lines {
		lineWithoutPrefix := strings.Replace(line, "Card 1: ", "", 1)
		pair := strings.Split(lineWithoutPrefix, "|")
		winningNumbers, myNumbers := strings.Fields(pair[0]), strings.Fields(pair[1])
		card := Card{WinningNumbers: winningNumbers, MyNumbers: myNumbers, AmountOfCards: 1}
		cards = append(cards, card)
	}
	for i, card := range cards {
		wins := calculateNumberOfWins(card)
		for win := 0; win < wins; win++ {
			cards[i+1+win].AmountOfCards += card.AmountOfCards
		}
	}
	sumOfCards := 0
	for _, card := range cards {
		sumOfCards += card.AmountOfCards
	}
	return sumOfCards
}

func calculateNumberOfWins(card Card) int {
	wins := 0
	for _, myNumber := range card.MyNumbers {
		if slices.Contains(card.WinningNumbers, myNumber) {
			wins++
		}
	}
	return wins
}

func main() {
	day := "day04"
	exampleResult := solve1("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 13, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 30, exampleResult)
	result2 := solve2("2023/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
