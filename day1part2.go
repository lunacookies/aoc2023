package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("day1input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		os.Exit(1)
	}

	digits := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine",
	}

	trimmed := strings.Trim(string(file), "\n")
	lines := strings.Split(trimmed, "\n")
	sum := 0

	for _, line := range lines {
		var firstDigit, lastDigit int
		firstPosition := math.MaxInt
		lastPosition := -1

		for i, b := range line {
			isDigit := b >= '0' && b <= '9'
			if !isDigit {
				continue
			}

			digit := int(b - '0')
			if i < firstPosition {
				firstDigit = digit
				firstPosition = i
			}
			lastDigit = digit
			lastPosition = i
		}

		for digitIndex, digitText := range digits {
			digit := digitIndex + 1

			position := strings.Index(line, digitText)
			if position != -1 {
				if position < firstPosition {
					firstDigit = digit
					firstPosition = position
				}
			}

			position = strings.LastIndex(line, digitText)
			if position != -1 {
				if position > lastPosition {
					lastDigit = digit
					lastPosition = position
				}
			}
		}

		sum += firstDigit*10 + lastDigit
	}

	fmt.Println(sum)
}
