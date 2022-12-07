package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"unicode"
)

func solve1(inputPath string) int {
	input := IOUtils.ReadInputRunes(inputPath)
	priority := 0
	for _, rucksacks := range input {
		a, b := split(rucksacks)
		setA := make(map[rune]bool)
		for _, item := range a {
			setA[item] = true
		}
		for _, item := range b {
			if _, ok := setA[item]; ok {
				if unicode.IsUpper(item) {
					priority += int(item) - 38
				} else {
					priority += int(item) - 96
				}
				break
			}
		}
	}
	return priority
}

func solve2(inputPath string) int {
	input := IOUtils.ReadInputRunes(inputPath)
	priority := 0
	for i := 0; i < len(input); i += 3 {
		a := input[i]
		b := input[i+1]
		c := input[i+2]
		interSecAB := intersection(a, b)
		interSecAll := intersection(interSecAB, c)
		tag := interSecAll[0]
		if unicode.IsUpper(tag) {
			priority += int(tag) - 38
		} else {
			priority += int(tag) - 96
		}
	}
	return priority
}

func intersection(s1, s2 []rune) (inter []rune) {
	hash := make(map[rune]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

//Remove dups from slice.
func removeDups(elements []rune) (nodups []rune) {
	encountered := make(map[rune]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}
func split(rucksacks []rune) (a []rune, b []rune) {
	a = rucksacks[:len(rucksacks)/2]
	b = rucksacks[len(rucksacks)/2:]
	return a, b
}

func main() {
	day := "day03"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 157, exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 70, exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
