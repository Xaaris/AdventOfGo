package main

import (
	"AdventOfGo/IOUtils"
	"fmt"
)

func solve1(inputPath string) int {
	f := IOUtils.ReadInputDigits(inputPath)
	v := make([][]bool, len(f))
	for i := range v {
		v[i] = make([]bool, len(f[0]))
	}
	for i := range v {
		for j := range v[0] {
			if i == 0 || j == 0 || i == len(v)-1 || j == len(v[0])-1 {
				v[i][j] = true
				fmt.Print("I")
			} else if isBiggerLeft(i, j, f) || isBiggerRight(i, j, f) || isBiggerUp(i, j, f) || isBiggerDown(i, j, f) {
				v[i][j] = true
				fmt.Print("I")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	visibleCounter := 0
	for i := range v {
		for j := range v[0] {
			if v[i][j] == true {
				visibleCounter++
			}
		}
	}
	return visibleCounter
}

func isBiggerLeft(i, j int, f [][]int) bool {
	for _, h := range f[i][0:j] {
		if h >= f[i][j] {
			return false
		}
	}
	return true
}

func isBiggerRight(i, j int, f [][]int) bool {
	for _, h := range f[i][j+1 : len(f)] {
		if h >= f[i][j] {
			return false
		}
	}
	return true
}

func isBiggerUp(i, j int, f [][]int) bool {
	for y := 0; y < i; y++ {
		if f[y][j] >= f[i][j] {
			return false
		}
	}
	return true
}

func isBiggerDown(i, j int, f [][]int) bool {
	for y := i + 1; y < len(f[0]); y++ {
		if f[y][j] >= f[i][j] {
			return false
		}
	}
	return true
}

func solve2(inputPath string) int {
	f := IOUtils.ReadInputDigits(inputPath)
	v := make([][]int, len(f))
	for i := range v {
		v[i] = make([]int, len(f[0]))
	}
	for i := range v {
		for j := range v[0] {
			if i == 0 || j == 0 || i == len(v)-1 || j == len(v[0])-1 {
				v[i][j] = 0
			} else {
				v[i][j] = viewRangeLeft(i, j, f) * viewRangeRight(i, j, f) * viewRangeUp(i, j, f) * viewRangeDown(i, j, f)
			}
		}
	}

	maxScenicScore := 0
	for i := range v {
		for j := range v[0] {
			if v[i][j] > maxScenicScore {
				maxScenicScore = v[i][j]
			}
		}
	}

	return maxScenicScore
}

func viewRangeLeft(i, j int, f [][]int) int {
	viewRange := 0
	for x := j - 1; x >= 0; x-- {
		viewRange++
		if f[i][x] >= f[i][j] {
			break
		}
	}
	return viewRange
}

func viewRangeRight(i, j int, f [][]int) int {
	viewRange := 0
	for x := j + 1; x < len(f); x++ {
		viewRange++
		if f[i][x] >= f[i][j] {
			break
		}
	}
	return viewRange
}

func viewRangeUp(i, j int, f [][]int) int {
	viewRange := 0
	for y := i - 1; y >= 0; y-- {
		viewRange++
		if f[y][j] >= f[i][j] {
			break
		}
	}
	return viewRange
}

func viewRangeDown(i, j int, f [][]int) int {
	viewRange := 0
	for y := i + 1; y < len(f[0]); y++ {
		viewRange++
		if f[y][j] >= f[i][j] {
			break
		}
	}
	return viewRange
}

func main() {
	day := "day08"
	exampleResult := solve1("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 21, exampleResult)
	result1 := solve1("2022/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

	exampleResult = solve2("2022/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 8, exampleResult)
	result2 := solve2("2022/" + day + "/data.txt")
	fmt.Printf("Result 2: %d\n", result2)
}
