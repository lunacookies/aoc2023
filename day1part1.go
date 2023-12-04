package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("day1input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		os.Exit(1)
	}

	trimmed := strings.Trim(string(file), "\n")
	lines := strings.Split(trimmed, "\n")
	sum := 0

	for _, line := range lines {
		var first, last int

		for _, b := range line {
			isDigit := b >= '0' && b <= '9'
			if !isDigit {
				continue
			}

			digit := int(b - '0')
			if first == 0 {
				first = digit
			}
			last = digit
		}

		sum += first*10 + last
	}

	fmt.Println(sum)
}
