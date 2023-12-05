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
}

var myCards []scratchCard

func parseCards() {
	var data []byte
	var err error
	var card *scratchCard

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
		// get my numbers
		for _, numStr := range strings.Split(strings.TrimSpace(temps[1]), " ") {
			if numStr == "" {
				continue
			}
			if num, err = strconv.Atoi(numStr); err != nil {
				panic(err)
			}
			card.myNums = append(card.myNums, num)
		}
		myCards = append(myCards, *card)
	}
}

func part1() {
	var total = 0
	for _, card := range myCards {
		var points = 0
		for _, num := range card.winningNums {
			if slices.Contains(card.myNums, num) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		total += points
	}
	fmt.Printf("P1 Total: %d\n", total)
}

func part2() {

}

func main() {
	parseCards()
	part1()
	part2()
}
