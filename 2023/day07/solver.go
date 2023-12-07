package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"AdventOfGo/IOUtils"
)

type hand struct {
	cards        [5]int
	typeStrength int
	bid          int
}

func (h hand) String() string {
	cards := []string{}
	for _, card := range h.cards {
		switch card {
		case 14:
			cards = append(cards, "A")
		case 13:
			cards = append(cards, "K")
		case 12:
			cards = append(cards, "Q")
		case 11, 1:
			cards = append(cards, "J")
		case 10:
			cards = append(cards, "T")
		default:
			cards = append(cards, strconv.Itoa(card))
		}
	}
	typeString := ""
	switch h.typeStrength {
	case 6:
		typeString = "5 of a kind"
	case 5:
		typeString = "4 of a kind"
	case 4:
		typeString = "Full house"
	case 3:
		typeString = "3 of a kind"
	case 2:
		typeString = "2 pair"
	case 1:
		typeString = "1 pair"
	default:
		typeString = "High card"
	}
	return fmt.Sprintf("%v (%v) bid: %v", cards, typeString, h.bid)
}

func solve1(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	var hands []hand
	for _, line := range lines {
		fields := strings.Fields(line)
		cards := stringToCards(fields[0])
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, hand{cards, getTypeStrength(cards), bid})
	}
	handSortFn := func(a, b hand) int {
		// Sort by type strength
		if n := cmp.Compare(a.typeStrength, b.typeStrength); n != 0 {
			return n
		}
		// Same type, sort by highest card
		for i := 0; i < len(a.cards); i++ {
			if n := cmp.Compare(a.cards[i], b.cards[i]); n != 0 {
				return n
			}
		}
		panic("Same hand")
	}
	slices.SortFunc(hands, handSortFn)
	fmt.Println("Sorted:")
	for i, hand := range hands {
		fmt.Printf("%v: %v\n", i, hand)
	}
	score := 0
	for i, hand := range hands {
		score += hand.bid * (i + 1)
	}
	return score
}

func stringToCards(cardString string) [5]int {
	cards := []int{}
	runes := []rune(cardString)
	for _, r := range runes {
		switch r {
		case 'A':
			cards = append(cards, 14)
		case 'K':
			cards = append(cards, 13)
		case 'Q':
			cards = append(cards, 12)
		case 'J':
			cards = append(cards, 11)
		case 'T':
			cards = append(cards, 10)
		default:
			atoi, _ := strconv.Atoi(string(r))
			cards = append(cards, atoi)
		}
	}
	return [5]int(cards)
}

func stringToCardsWithJoker(cardString string) [5]int {
	cards := []int{}
	runes := []rune(cardString)
	for _, r := range runes {
		switch r {
		case 'A':
			cards = append(cards, 14)
		case 'K':
			cards = append(cards, 13)
		case 'Q':
			cards = append(cards, 12)
		case 'J':
			cards = append(cards, 1)
		case 'T':
			cards = append(cards, 10)
		default:
			atoi, _ := strconv.Atoi(string(r))
			cards = append(cards, atoi)
		}
	}
	return [5]int(cards)
}

func getTypeStrength(cards [5]int) int {
	switch {
	case is5OfAKind(cards):
		return 6
	case is4OfAKind(cards):
		return 5
	case isFullHouse(cards):
		return 4
	case is3OfAKind(cards):
		return 3
	case isTwoPair(cards):
		return 2
	case isOnePair(cards):
		return 1
	default:
		return 0
	}
}

func getTypeStrengthWithJoker(cards [5]int) int {
	switch {
	case is5OfAKindWithJoker(cards):
		return 6
	case is4OfAKindWithJoker(cards):
		return 5
	case isFullHouseWithJoker(cards):
		return 4
	case is3OfAKindWithJoker(cards):
		return 3
	case isTwoPairWithJoker(cards):
		return 2
	case isOnePairWithJoker(cards):
		return 1
	default:
		return 0
	}
}

func is5OfAKind(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	if len(mapOfCards) == 1 {
		return true
	}
	return false
}

func is5OfAKindWithJoker(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	jokers := mapOfCards[1]
	if jokers == 5 {
		return true
	}
	for k, v := range mapOfCards {
		if k == 1 {
			continue
		}
		if v+jokers == 5 {
			return true
		}
	}
	return false
}

func is4OfAKind(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	if len(mapOfCards) != 2 {
		return false
	}
	for _, v := range mapOfCards {
		if v == 4 {
			return true
		}
	}
	return false
}

func is4OfAKindWithJoker(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	jokers := mapOfCards[1]
	for k, v := range mapOfCards {
		if k == 1 {
			continue
		}
		if v+jokers == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	if len(mapOfCards) != 2 {
		return false
	}
	for _, v := range mapOfCards {
		if v == 3 {
			return true
		}
	}
	return false
}

func isFullHouseWithJoker(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	jokers := mapOfCards[1]
	if len(mapOfCards) == 2 {
		for k, v := range mapOfCards {
			if k == 1 {
				continue
			}
			if v == 3 {
				return true
			}
		}
	}
	if len(mapOfCards) == 3 && jokers > 0 {
		for _, v := range mapOfCards {
			if v+jokers == 3 {
				return true
			}
		}
	}
	return false
}

func is3OfAKind(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	if len(mapOfCards) != 3 {
		return false
	}
	for _, v := range mapOfCards {
		if v == 3 {
			return true
		}
	}
	return false
}

func is3OfAKindWithJoker(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	jokers := mapOfCards[1]
	for k, v := range mapOfCards {
		if k == 1 {
			continue
		}
		if v+jokers == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	if len(mapOfCards) != 3 {
		return false
	}
	for _, v := range mapOfCards {
		if v == 1 {
			return true
		}
	}
	return false
}

func isTwoPairWithJoker(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	jokers := mapOfCards[1]
	if jokers == 2 {
		return true
	}
	if jokers == 1 {
		for k, v := range mapOfCards {
			if k == 1 {
				continue
			}
			if v == 2 {
				return true
			}
		}
	}
	if len(mapOfCards) == 3 {
		for _, v := range mapOfCards {
			if v == 1 {
				return true
			}
		}
	}
	return false
}

func isOnePair(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	if len(mapOfCards) == 4 {
		return true
	}
	return false
}

func isOnePairWithJoker(cards [5]int) bool {
	mapOfCards := toMapOfCards(cards)
	if len(mapOfCards) == 4 && mapOfCards[1] == 0 {
		return true
	}
	if len(mapOfCards) == 5 && mapOfCards[1] == 1 {
		return true
	}
	return false
}

func toMapOfCards(cards [5]int) map[int]int {
	mapOfCards := make(map[int]int)
	for _, card := range cards {
		mapOfCards[card]++
	}
	return mapOfCards
}

func solve2(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	var hands []hand
	for _, line := range lines {
		fields := strings.Fields(line)
		cards := stringToCardsWithJoker(fields[0])
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, hand{cards, getTypeStrengthWithJoker(cards), bid})
	}
	handSortFn := func(a, b hand) int {
		// Sort by type strength
		if n := cmp.Compare(a.typeStrength, b.typeStrength); n != 0 {
			return n
		}
		// Same type, sort by highest card
		for i := 0; i < len(a.cards); i++ {
			if n := cmp.Compare(a.cards[i], b.cards[i]); n != 0 {
				return n
			}
		}
		panic("Same hand")
	}
	slices.SortFunc(hands, handSortFn)
	fmt.Println("Sorted:")
	for i, hand := range hands {
		fmt.Printf("%v: %v\n", i, hand)
	}
	score := 0
	for i, hand := range hands {
		score += hand.bid * (i + 1)
	}
	return score
}

func main() {
	day := "day07"
	exampleResult := solve1("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 6440, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 5905, exampleResult)
	result2 := solve2("2023/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
