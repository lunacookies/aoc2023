package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day2input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		return
	}

	trimmed := strings.Trim(string(file), "\n")
	games := strings.Split(trimmed, "\n")
	sum := 0

	for _, game := range games {
		colorCounts := make(map[string]int)
		_, sets, _ := strings.Cut(game, ": ")

		for _, set := range strings.Split(sets, "; ") {
			for _, countAndColor := range strings.Split(set, ", ") {
				countString, color, _ := strings.Cut(countAndColor, " ")
				count, _ := strconv.Atoi(countString)

				if colorCounts[color] < count {
					colorCounts[color] = count
				}
			}
		}

		sum += colorCounts["red"] * colorCounts["green"] * colorCounts["blue"]
	}

	fmt.Println(sum)
}
