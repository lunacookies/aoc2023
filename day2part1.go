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
		os.Exit(1)
	}

	trimmed := strings.Trim(string(file), "\n")
	games := strings.Split(trimmed, "\n")
	idSum := 0

	for i, game := range games {
		if isGameValid(game) {
			gameID := i + 1
			idSum += gameID
		}
	}

	fmt.Println(idSum)
}

func isGameValid(game string) bool {
	_, sets, _ := strings.Cut(game, ": ")

	for _, set := range strings.Split(sets, "; ") {
		if !isSetValid(set) {
			return false
		}
	}

	return true

}

func isSetValid(set string) bool {
	maxCountForColor := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, countAndColor := range strings.Split(set, ", ") {
		countString, color, _ := strings.Cut(countAndColor, " ")
		count, _ := strconv.Atoi(countString)

		if count > maxCountForColor[color] {
			return false
		}
	}

	return true
}
