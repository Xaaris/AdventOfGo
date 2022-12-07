package IOUtils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInputStrings(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadInputInts(path string) []int {
	return mapToInt(ReadInputStrings(path))
}

func ReadInputRunes(path string) [][]rune {
	return mapToRunes(ReadInputStrings(path))
}

func mapToInt(strings []string) []int {
	ints := make([]int, len(strings))

	var err error
	for i, s := range strings {
		ints[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
	}

	return ints
}

func mapToRunes(strings []string) [][]rune {
	runes := make([][]rune, len(strings))

	for i, s := range strings {
		runes[i] = []rune(s)
	}

	return runes
}
