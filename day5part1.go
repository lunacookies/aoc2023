package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type category string

type almanacMap struct {
	destination   category
	almanacRanges []almanacRange
}

type almanacRange struct {
	destinationStart, sourceStart, length int
}

func main() {
	file, err := os.ReadFile("day5input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		os.Exit(1)
	}

	trimmed := strings.Trim(string(file), "\n")
	sections := strings.Split(trimmed, "\n\n")

	seedsS, _ := strings.CutPrefix(sections[0], "seeds: ")
	seeds := make([]int, 0)

	for _, seedS := range strings.Split(seedsS, " ") {
		seed, _ := strconv.Atoi(seedS)
		seeds = append(seeds, seed)
	}

	almanacMaps := make(map[category]almanacMap)

	for _, section := range sections[1:] {
		header, body, _ := strings.Cut(section, " map:\n")

		sourceS, destinationS, _ := strings.Cut(header, "-to-")
		source := category(sourceS)
		destination := category(destinationS)
		almanacRanges := make([]almanacRange, 0)

		for _, line := range strings.Split(body, "\n") {
			components := strings.Split(line, " ")
			destinationStart, _ := strconv.Atoi(components[0])
			sourceStart, _ := strconv.Atoi(components[1])
			length, _ := strconv.Atoi(components[2])

			r := almanacRange{destinationStart, sourceStart, length}
			almanacRanges = append(almanacRanges, r)
		}

		almanacMaps[source] = almanacMap{destination, almanacRanges}
	}

	lowest := math.MaxInt

	for _, seed := range seeds {
		currentCategory := category("seed")
		currentValue := seed

		for {
			almanacMap, ok := almanacMaps[currentCategory]
			if !ok {
				break
			}

			for _, almanacRange := range almanacMap.almanacRanges {
				start := almanacRange.sourceStart
				end := almanacRange.sourceStart + almanacRange.length - 1
				inRange := currentValue >= start && currentValue <= end
				if !inRange {
					continue
				}

				offset := currentValue - start
				currentValue = almanacRange.destinationStart + offset
				break
			}

			currentCategory = almanacMap.destination
		}

		if currentValue < lowest {
			lowest = currentValue
		}
	}

	fmt.Println(lowest)
}
