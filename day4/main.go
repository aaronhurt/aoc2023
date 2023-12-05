package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type scratchCard struct {
	id          int
	winningNums []int
	myNums      []int
	points      int
}
type scratchCards struct {
	cards []scratchCard
	total int
}

func parseCards() scratchCards {
	var data []byte
	var err error
	var card *scratchCard
	var cards scratchCards

	if data, err = os.ReadFile("./input.txt"); err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}
		card = new(scratchCard)

		// get card number
		if _, err = fmt.Sscanf(line, "Card\t%d:", &card.id); err != nil {
			panic(err)
		}

		// split numbers
		temps := strings.Split(line[strings.Index(line, ":")+1:], "|")
		var num = 0

		// get winning numbers
		for _, numStr := range strings.Split(strings.TrimSpace(temps[0]), " ") {
			if numStr == "" {
				continue
			}
			if num, err = strconv.Atoi(numStr); err != nil {
				panic(err)
			}
			card.winningNums = append(card.winningNums, num)
		}
		// get my numbers and calculate points
		for _, numStr := range strings.Split(strings.TrimSpace(temps[1]), " ") {
			if numStr == "" {
				continue
			}
			if num, err = strconv.Atoi(numStr); err != nil {
				panic(err)
			}
			card.myNums = append(card.myNums, num)
			if slices.Contains(card.winningNums, num) {
				if card.points == 0 {
					card.points = 1
				} else {
					card.points *= 2
				}
			}
		}

		// add card and total points
		cards.cards = append(cards.cards, *card)
		cards.total += card.points
	}
	return cards
}

func part1() {
	cards := parseCards()
	fmt.Printf("P1 Total: %d\n", cards.total)
}

func part2() {

}

func main() {
	part1()
	part2()
}
