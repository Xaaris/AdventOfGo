package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"AdventOfGo/IOUtils"
)

type rangeMapping struct {
	fromStart int
	fromEnd   int
	delta     int
}

func solve1(inputPath string) int {
	lines := IOUtils.ReadInputStrings(inputPath)
	seeds := []int{}
	seedToSoil := []rangeMapping{}
	soilToFertilizer := []rangeMapping{}
	fertilizerToWater := []rangeMapping{}
	waterToLight := []rangeMapping{}
	lightToTemperature := []rangeMapping{}
	temperatureToHumidity := []rangeMapping{}
	humidityToLocation := []rangeMapping{}
	for i, line := range lines {
		if strings.HasPrefix(line, "seeds: ") {
			seeds = stringsToInts(strings.Fields(line[7:]))
			continue
		}
		if line == "" {
			continue
		}
		if line == "seed-to-soil map:" {
			seedToSoil = readRangeMapping(lines[i+1:])
		}
		if line == "soil-to-fertilizer map:" {
			soilToFertilizer = readRangeMapping(lines[i+1:])
		}
		if line == "fertilizer-to-water map:" {
			fertilizerToWater = readRangeMapping(lines[i+1:])
		}
		if line == "water-to-light map:" {
			waterToLight = readRangeMapping(lines[i+1:])
		}
		if line == "light-to-temperature map:" {
			lightToTemperature = readRangeMapping(lines[i+1:])
		}
		if line == "temperature-to-humidity map:" {
			temperatureToHumidity = readRangeMapping(lines[i+1:])
		}
		if line == "humidity-to-location map:" {
			humidityToLocation = readRangeMapping(lines[i+1:])
		}
	}
	fmt.Printf(
		"seeds: %v\n"+
			"seedToSoil: %v\n"+
			"soilToFertilizer: %v\n"+
			"fertilizerToWater: %v\n"+
			"waterToLight: %v\n"+
			"lightToTemperature: %v\n"+
			"temperatureToHumidity: %v\n"+
			"humidityToLocation: %v\n",
		seeds,
		seedToSoil,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation,
	)
	locations := []int{}
	for _, seed := range seeds {
		soil := seed
		for _, mapping := range seedToSoil {
			if seed >= mapping.fromStart && seed <= mapping.fromEnd {
				soil += mapping.delta
				break
			}
		}
		fertilizer := soil
		for _, mapping := range soilToFertilizer {
			if soil >= mapping.fromStart && soil <= mapping.fromEnd {
				fertilizer += mapping.delta
				break
			}
		}
		water := fertilizer
		for _, mapping := range fertilizerToWater {
			if fertilizer >= mapping.fromStart && fertilizer <= mapping.fromEnd {
				water += mapping.delta
				break
			}
		}
		light := water
		for _, mapping := range waterToLight {
			if water >= mapping.fromStart && water <= mapping.fromEnd {
				light += mapping.delta
				break
			}
		}
		temperature := light
		for _, mapping := range lightToTemperature {
			if light >= mapping.fromStart && light <= mapping.fromEnd {
				temperature += mapping.delta
				break
			}
		}
		humidity := temperature
		for _, mapping := range temperatureToHumidity {
			if temperature >= mapping.fromStart && temperature <= mapping.fromEnd {
				humidity += mapping.delta
				break
			}
		}
		location := humidity
		for _, mapping := range humidityToLocation {
			if humidity >= mapping.fromStart && humidity <= mapping.fromEnd {
				location += mapping.delta
				break
			}
		}
		locations = append(locations, location)
	}
	return slices.Min(locations)
}

func stringsToInts(strings []string) []int {
	ints := []int{}
	for _, str := range strings {
		atoi, _ := strconv.Atoi(str)
		ints = append(ints, atoi)
	}
	return ints
}

func readRangeMapping(lines []string) []rangeMapping {
	mappings := []rangeMapping{}
	for _, line := range lines {
		if line == "" {
			return mappings
		}
		ranges := stringsToInts(strings.Fields(line))

		mappings = append(mappings, rangeMapping{
			fromStart: ranges[1],
			fromEnd:   ranges[1] + ranges[2] - 1,
			delta:     ranges[0] - ranges[1],
		})
	}
	return mappings
}

func main() {
	day := "day05"
	exampleResult := solve1("2023/" + day + "/example.txt")
	fmt.Printf("expected: %d, got %d\n", 35, exampleResult)
	result1 := solve1("2023/" + day + "/data.txt")
	fmt.Printf("Result 1: %d\n", result1)

}
