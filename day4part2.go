package main

import (
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type numberSet struct {
	lo, hi uint64
}

func (s *numberSet) insert(n int) {
	if n > 63 {
		s.hi |= 1 << (n - 64)
	} else {
		s.lo |= 1 << n
	}
}

func (s *numberSet) insertSpaceSeparated(str string) {
	for _, ns := range strings.Split(str, " ") {
		if len(ns) == 0 {
			continue
		}

		n, _ := strconv.Atoi(ns)
		s.insert(n)
	}
}

func (s1 numberSet) intersection(s2 numberSet) numberSet {
	return numberSet{lo: s1.lo & s2.lo, hi: s1.hi & s2.hi}
}

func (s numberSet) count() int {
	return bits.OnesCount64(s.lo) + bits.OnesCount64(s.hi)
}

type card struct {
	winningNumbers numberSet
	myNumbers      numberSet
}

func main() {
	file, err := os.ReadFile("day4input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		os.Exit(1)
	}

	trimmed := strings.Trim(string(file), "\n")
	lines := strings.Split(trimmed, "\n")
	cards := make([]card, len(lines))

	for i, card := range lines {
		_, card, _ := strings.Cut(card, ": ")
		winningNumbers, myNumbers, _ := strings.Cut(card, " | ")

		cards[i].winningNumbers.insertSpaceSeparated(winningNumbers)
		cards[i].myNumbers.insertSpaceSeparated(myNumbers)
	}

	cardsToProcess := make([]int, 0, len(cards))
	for i := range cards {
		cardsToProcess = append(cardsToProcess, i)
	}

	processedCount := 0

	for len(cardsToProcess) > 0 {
		processedCount++
		cardID := cardsToProcess[len(cardsToProcess)-1]
		card := cards[cardID]
		cardsToProcess = cardsToProcess[:len(cardsToProcess)-1]

		matchCount := card.winningNumbers.intersection(card.myNumbers).count()

		if matchCount == 0 {
			continue
		}

		for id := cardID + 1; id <= cardID+matchCount; id++ {
			cardsToProcess = append(cardsToProcess, id)
		}

	}

	fmt.Println(processedCount)
}
