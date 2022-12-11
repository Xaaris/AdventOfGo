package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	parent   *node
	subNodes []*node
}

func solve1(inputPath string) int {
	input := IOUtils.ReadInputStrings(inputPath)

	root := node{name: "/"}
	var currentNode *node
	for _, line := range input {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd ..") {
				currentNode = currentNode.parent
			} else if line[2:4] == "cd" {
				directory := strings.TrimPrefix(line, "$ cd ")
				if directory == "/" {
					currentNode = &root
				} else {
					for _, subNode := range currentNode.subNodes {
						if subNode.name == directory {
							currentNode = subNode
							break
						}
					}
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			newNode := node{name: strings.TrimPrefix(line, "dir "), parent: currentNode}
			currentNode.subNodes = append(currentNode.subNodes, &newNode)
		} else {
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			currentNode.size += size
		}
	}

	root.calculateTotalDirSizes()

	listOfNodes := root.flatten()
	totalSize := 0
	for _, node := range listOfNodes {
		if node.size <= 100000 {
			totalSize += node.size
		}
	}

	return totalSize
}

func solve2(inputPath string) int {

	input := IOUtils.ReadInputStrings(inputPath)

	root := node{name: "/"}
	var currentNode *node
	for _, line := range input {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd ..") {
				currentNode = currentNode.parent
			} else if line[2:4] == "cd" {
				directory := strings.TrimPrefix(line, "$ cd ")
				if directory == "/" {
					currentNode = &root
				} else {
					for _, subNode := range currentNode.subNodes {
						if subNode.name == directory {
							currentNode = subNode
							break
						}
					}
				}
			}
		} else if strings.HasPrefix(line, "dir") {
			newNode := node{name: strings.TrimPrefix(line, "dir "), parent: currentNode}
			currentNode.subNodes = append(currentNode.subNodes, &newNode)
		} else {
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			currentNode.size += size
		}
	}

	root.calculateTotalDirSizes()

	listOfNodes := root.flatten()
	sort.Slice(listOfNodes, func(i, j int) bool {
		return listOfNodes[i].size < listOfNodes[j].size
	})
	freeSpace := 70000000 - root.size
	spaceNeeded := 30000000 - freeSpace

	for _, node := range listOfNodes {
		if node.size >= spaceNeeded {
			return node.size
		}
	}
	return 0
}

func (n *node) calculateTotalDirSizes() int {
	if len(n.subNodes) == 0 {
		return n.size
	} else {
		for _, subNode := range n.subNodes {
			n.size += subNode.calculateTotalDirSizes()
		}
		return n.size
	}
}

func (n *node) flatten() []*node {
	if len(n.subNodes) == 0 {
		return []*node{n}
	} else {
		var list []*node
		for _, subNode := range n.subNodes {
			list = append(list, subNode.flatten()...)
		}
		return append(list, n)
	}
}

func (n *node) sumUpSmallDirs() int {
	if len(n.subNodes) == 0 {
		if n.size <= 100000 {
			return n.size
		} else {
			return 0
		}
	} else {
		for _, subNode := range n.subNodes {
			if n.size <= 100000 {

			}
			n.size += subNode.sumUpSmallDirs()
		}
		return n.size
	}
}

func main() {
	day := "day07"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 95437, exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 24933642, exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
