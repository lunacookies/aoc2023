package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type schematic []string

func main() {
	file, err := os.ReadFile("day3input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		os.Exit(1)
	}

	trimmed := strings.Trim(string(file), "\n")
	s := schematic(strings.Split(trimmed, "\n"))
	sum := 0

	for y := range s {
		for x := range s[y] {
			atStartOfNumber := s.isDigitAt(x, y) && !s.isDigitAt(x-1, y)
			if !atStartOfNumber {
				continue
			}

			length := 0
			for ; s.isDigitAt(x+length, y); length++ {
			}

			if s.isPartNumberAt(x, y, length) {
				partNumber, _ := strconv.Atoi(s.region(x, y, length))
				sum += partNumber
			}
		}
	}

	fmt.Println(sum)
}

func (s schematic) isPartNumberAt(x, y, length int) bool {
	for yCand := y - 1; yCand <= y+1; yCand++ {
		for xCand := x - 1; xCand <= x+length; xCand++ {
			if s.isSymbolAt(xCand, yCand) {
				return true
			}
		}
	}

	return false
}

func (s schematic) isSymbolAt(x, y int) bool {
	if !s.inBounds(x, y) {
		return false
	}

	return isSymbol(s.at(x, y))
}

func (s schematic) isDigitAt(x, y int) bool {
	if !s.inBounds(x, y) {
		return false
	}

	return isDigit(s.at(x, y))
}

func (s schematic) inBounds(x, y int) bool {
	return x >= 0 && x < s.width() && y >= 0 && y < s.height()
}

func (s schematic) at(x, y int) byte {
	return s[y][x]
}

func (s schematic) region(x, y, length int) string {
	return s[y][x : x+length]
}

func (s schematic) width() int {
	return len(s[0])
}

func (s schematic) height() int {
	return len(s)
}

func isSymbol(b byte) bool {
	return !isDigit(b) && b != '.'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
