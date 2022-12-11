// This file has been generated with the help of chatGPT.
//Only minor fixes have been done to the version generated by chatGPT to fix a small mistake.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Grid2 represents the grid on which the head and tail move
type Grid2 struct {
	snake [10]*Pos
}

// NewGrid2 creates a new grid with the given head and tail positions
func NewGrid2() *Grid2 {
	var snake [10]*Pos
	for i := 0; i < len(snake); i++ {
		snake[i] = &Pos{0, 0}
	}
	return &Grid2{snake}
}

// Update2 updates the positions of the head and tail according to the given move
func (g *Grid2) Update2(move string, positions map[Pos]struct{}) {
	head := g.snake[0]

	// Update the position of the head
	switch move {
	case "U":
		head.move(0, -1)
	case "D":
		head.move(0, 1)
	case "L":
		head.move(-1, 0)
	case "R":
		head.move(1, 0)
	}

	for i := 1; i < len(g.snake); i++ {
		head := g.snake[i-1]
		tail := g.snake[i]
		// Update the position of the tail
		if tail.isAdjacent(head) {
			// If the head and tail are touching, the tail does not move
			return
		} else if tail.isInSameRow(head) {
			// If the head is two steps directly up, down, left, or right from the tail,
			// the tail moves one step in that direction
			if head.x == tail.x+2 {
				tail.move(1, 0)
			} else if head.x == tail.x-2 {
				tail.move(-1, 0)
			}
		} else if tail.isInSameCol(head) {
			if head.y == tail.y+2 {
				tail.move(0, 1)
			} else if head.y == tail.y-2 {
				tail.move(0, -1)
			}
		} else {
			// Otherwise, if the head and tail are not in the same row or column,
			// the tail moves one step diagonally to keep up with the head
			if head.x > tail.x {
				tail.move(1, 0)
			} else {
				tail.move(-1, 0)
			}
			if head.y > tail.y {
				tail.move(0, 1)
			} else {
				tail.move(0, -1)
			}
		}
	}

	// Mark the new position of the tail as visited
	positions[*g.snake[9]] = struct{}{}
}

func main() {
	// Open the file
	file, err := os.Open("2022/day09/data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the input from the file
	scanner := bufio.NewScanner(file)

	// Create a new grid with the head and tail starting at the same position
	g := NewGrid2()

	// Create a map to keep track of the positions visited by the tail
	positions := map[Pos]struct{}{}
	positions[Pos{0, 0}] = struct{}{}

	// Update the positions of the head and tail according to the moves
	for scanner.Scan() {
		input := scanner.Text()
		moves := strings.Split(input, " ")
		repetitions, _ := strconv.Atoi(moves[1])
		for i := 0; i < repetitions; i++ {
			fmt.Println(moves)
			g.Update2(moves[0], positions)
		}
	}

	// Print the final number of places that the tail has visited
	fmt.Println(len(positions))

}