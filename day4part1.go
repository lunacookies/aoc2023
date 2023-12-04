package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("day4input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		os.Exit(1)
	}

	trimmed := strings.Trim(string(file), "\n")
	lines := strings.Split(trimmed, "\n")
	points := 0

	for _, card := range lines {
		_, card, _ := strings.Cut(card, ": ")
		winningNumbersString, myNumbersString, _ := strings.Cut(card, " | ")

		winningNumbers := make(map[string]bool)
		for _, winningNumber := range strings.Split(winningNumbersString, " ") {
			if len(winningNumber) == 0 {
				continue
			}

			winningNumbers[winningNumber] = true
		}

		matchCount := 0

		myNumbers := strings.Split(myNumbersString, " ")
		for _, myNumber := range myNumbers {
			if winningNumbers[myNumber] {
				matchCount++
			}
		}

		if matchCount >= 1 {
			points += 1 << (matchCount - 1)
		}
	}

	fmt.Println(points)
}
