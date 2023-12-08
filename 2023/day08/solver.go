package main

import (
	"fmt"
	"strings"

	"AdventOfGo/IOUtils"
)

func solve1(inputPath string) int {
	instructionsLR := IOUtils.ReadInputRunes(inputPath)[0]
	instructions := []int{}
	for _, instruction := range instructionsLR {
		if instruction == 'L' {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, 1)
		}
	}
	nodeStrings := IOUtils.ReadInputStrings(inputPath)[2:]
	nodes := map[string][2]string{}
	for _, nodeString := range nodeStrings {
		start := nodeString[0:3]
		left := nodeString[7:10]
		right := nodeString[12:15]
		nodes[start] = [2]string{left, right}
	}
	steps := 0
	currentNode := "AAA"
	for i := 0; true; i = (i + 1) % len(instructions) {
		if currentNode == "ZZZ" {
			return steps
		}
		currentNode = nodes[currentNode][instructions[i]]
		steps++
	}
	panic("no solution found")
}

func solve2(inputPath string) uint64 {
	instructionsLR := IOUtils.ReadInputRunes(inputPath)[0]
	instructions := []int{}
	for _, instruction := range instructionsLR {
		if instruction == 'L' {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, 1)
		}
	}
	nodeStrings := IOUtils.ReadInputStrings(inputPath)[2:]
	nodes := map[string][2]string{}
	for _, nodeString := range nodeStrings {
		start := nodeString[0:3]
		left := nodeString[7:10]
		right := nodeString[12:15]
		nodes[start] = [2]string{left, right}
	}
	steps := 0
	currentNodes := []string{}
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			currentNodes = append(currentNodes, k)
		}
	}
	zsFoundAt := map[int][]int{}
	for i := 0; true; i = (i + 1) % len(instructions) {
		for j, node := range currentNodes {
			if !strings.HasSuffix(node, "Z") {
				break
			}
			if j == len(currentNodes)-1 {
				return uint64(steps)
			}
		}
		for j, node := range currentNodes {
			if strings.HasSuffix(node, "Z") {
				if len(zsFoundAt[j]) <= 10 {
					zsFoundAt[j] = append(zsFoundAt[j], steps)
				}
			}
		}
		//if len(zsFoundAt) == len(currentNodes) {
		//	fmt.Println(zsFoundAt)
		//}
		newNodes := []string{}
		for _, node := range currentNodes {
			newNodes = append(newNodes, nodes[node][instructions[i]])
		}
		currentNodes = newNodes
		steps++

		if foundEnoughExamples(zsFoundAt, len(currentNodes)) {
			break
		}
	}
	deltas := map[int][]int{}
	for k, v := range zsFoundAt {
		for i := 1; i < len(v); i++ {
			deltas[k] = append(deltas[k], v[i]-v[i-1])
		}
	}
	fmt.Printf("deltas: %v\n", deltas)
	var estimatedSteps uint64 = 1
	for _, v := range deltas {
		estimatedSteps *= uint64(v[0])
	}

	// From here took the lcm via wolfram alpha which turned out to be 22289513667691
	return estimatedSteps
}

func foundEnoughExamples(zsFoundAt map[int][]int, numberOfStartNodes int) bool {
	if len(zsFoundAt) < numberOfStartNodes {
		return false
	}
	for _, v := range zsFoundAt {
		if len(v) < 3 {
			return false
		}
	}
	return true

}

func main() {
	day := "day08"
	exampleResult := solve1("2023/" + day + "/example1.txt")
	fmt.Printf("expected: %d, got %d\n", 2, exampleResult)
	exampleResult = solve1("2023/" + day + "/example2.txt")
	fmt.Printf("expected: %d, got %d\n", 6, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult2 := solve2("2023/" + day + "/example3.txt")
	fmt.Printf("expected: %d, got %d\n", 6, exampleResult2)
	result2 := solve2("2023/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
