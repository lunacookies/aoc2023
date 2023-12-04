package main

import (
	"fmt"
	"os"
	"strings"
)

type card struct {
	winningNumbers map[string]bool
	myNumbers      []string
}

func main() {
	file, err := os.ReadFile("day4input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to read input\n")
		os.Exit(1)
	}

	trimmed := strings.Trim(string(file), "\n")
	lines := strings.Split(trimmed, "\n")
	cards := make([]card, 0)

	for _, cardString := range lines {
		_, cardString, _ := strings.Cut(cardString, ": ")
		winningNumbersString, myNumbersString, _ :=
			strings.Cut(cardString, " | ")

		card := card{winningNumbers: make(map[string]bool)}

		for _, winningNumber := range strings.Split(winningNumbersString, " ") {
			if len(winningNumber) == 0 {
				continue
			}

			card.winningNumbers[winningNumber] = true
		}

		card.myNumbers = strings.Split(myNumbersString, " ")
		cards = append(cards, card)
	}

	cardsToProcess := make([]int, 0)
	for i := range cards {
		cardsToProcess = append(cardsToProcess, i)
	}

	processedCount := 0

	for len(cardsToProcess) > 0 {
		processedCount++
		cardID := cardsToProcess[len(cardsToProcess)-1]
		card := cards[cardID]
		cardsToProcess = cardsToProcess[:len(cardsToProcess)-1]

		matchCount := 0
		for _, myNumber := range card.myNumbers {
			if card.winningNumbers[myNumber] {
				matchCount++
			}
		}

		if matchCount == 0 {
			continue
		}

		for id := cardID + 1; id <= cardID+matchCount; id++ {
			cardsToProcess = append(cardsToProcess, id)
		}

	}

	fmt.Println(processedCount)
}
