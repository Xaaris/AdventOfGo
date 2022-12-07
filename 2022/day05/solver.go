package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type move struct {
	amount int
	from   int
	to     int
}

func solve1(inputPath string) string {
	input := IOUtils.ReadInputStrings(inputPath)
	crates := map[int]string{}
	var moves []move

	for _, str := range input {
		if strings.HasPrefix(strings.TrimSpace(str), "[") {
			crates = parseCrates(str, crates)
		} else if strings.HasPrefix(strings.TrimSpace(str), "move") {
			moves = append(moves, parseMoves(str))
		}
	}

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			crates = executeMove(crates, move.from, move.to)
		}
	}

	result := ""
	for i := 1; i <= len(crates); i++ {
		result += crates[i][len(crates[i])-1:]
	}
	return result
}

func solve2(inputPath string) string {
	input := IOUtils.ReadInputStrings(inputPath)
	crates := map[int]string{}
	var moves []move

	for _, str := range input {
		if strings.HasPrefix(strings.TrimSpace(str), "[") {
			crates = parseCrates(str, crates)
		} else if strings.HasPrefix(strings.TrimSpace(str), "move") {
			moves = append(moves, parseMoves(str))
		}
	}

	for _, move := range moves {
		crates = executeMoveCombo(crates, move.amount, move.from, move.to)
	}

	result := ""
	for i := 1; i <= len(crates); i++ {
		result += crates[i][len(crates[i])-1:]
	}
	return result
}

func parseCrates(str string, crates map[int]string) map[int]string {
	for i := 0; i < len(str); i += 4 {
		if str[i+1] == ' ' {
			continue
		}
		lane := i/4 + 1
		crates[lane] = string(str[i+1]) + crates[lane]
	}
	return crates
}

func parseMoves(str string) move {
	var exp = regexp.MustCompile(`.+\s(\d+)\s.+\s(\d+)\s.+\s(\d+)`)
	matches := exp.FindStringSubmatch(str)
	amount, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])
	return move{amount, from, to}
}

func executeMove(crates map[int]string, from int, to int) map[int]string {
	crates[to] = crates[to] + crates[from][len(crates[from])-1:]
	crates[from] = crates[from][:len(crates[from])-1]
	return crates
}

func executeMoveCombo(crates map[int]string, amount int, from int, to int) map[int]string {
	crates[to] = crates[to] + crates[from][len(crates[from])-amount:]
	crates[from] = crates[from][:len(crates[from])-amount]
	return crates
}

func main() {
	day := "day05"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %s, got %s\n", "CMZ", exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %s\n", result1)

	exampleResult = solve2("2022/" + day + "/example.txt")
	fmt.Printf("expected: %s, got %s\n", "MCD", exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %s\n", result2)
}
