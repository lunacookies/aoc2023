package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type seedRange struct {
	start, length int
}

type almanacMap struct {
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
	seedsSplit := strings.Split(seedsS, " ")
	seedRanges := make([]seedRange, 0)

	for i := 0; i < len(seedsSplit); i += 2 {
		start, _ := strconv.Atoi(seedsSplit[i])
		length, _ := strconv.Atoi(seedsSplit[i+1])
		seedRanges = append(seedRanges, seedRange{start, length})
	}

	almanacMaps := make([]almanacMap, 1)

	for _, section := range sections[1:] {
		_, body, _ := strings.Cut(section, " map:\n")

		almanacRanges := make([]almanacRange, 0)

		for _, line := range strings.Split(body, "\n") {
			components := strings.Split(line, " ")
			destinationStart, _ := strconv.Atoi(components[0])
			sourceStart, _ := strconv.Atoi(components[1])
			length, _ := strconv.Atoi(components[2])

			r := almanacRange{destinationStart, sourceStart, length}
			almanacRanges = append(almanacRanges, r)
		}

		m := almanacMap{almanacRanges}
		almanacMaps = append(almanacMaps, m)
	}

	total := 0
	for _, seedRange := range seedRanges {
		total += seedRange.length
	}

	lowestValues := make(chan int)

	for _, seedRange := range seedRanges {
		r := seedRange
		go func() {
			fmt.Printf("kicked off %+v\n", r)

			lowest := math.MaxInt

			for i := 0; i < r.length; i++ {
				seed := r.start + i
				currentValue := seed

				for _, almanacMap := range almanacMaps {
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
				}

				if currentValue < lowest {
					lowest = currentValue
				}
			}

			fmt.Printf("completed %+v with %d\n", r, lowest)
			lowestValues <- lowest
		}()
	}

	lowest := math.MaxInt
	i := 0
	for l := range lowestValues {
		fmt.Printf("got %d\n", l)
		if l < lowest {
			lowest = l
		}

		if i == len(seedRanges)-1 {
			break
		}

		i++
	}

	fmt.Println(lowest)
}
