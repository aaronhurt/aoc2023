package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	p1RedMax   = 12
	p1GreenMax = 13
	p1BlueMax  = 14
)

type drawData struct {
	red   int
	green int
	blue  int
}
type gameData struct {
	id    int
	draws []drawData
}

func readGameData() []gameData {
	var data []byte
	var err error
	var games []gameData

	if data, err = os.ReadFile("./input.txt"); err != nil {
		log.Fatal(err)
	}

	// split out game number and draws
	mainFieldsFunc := func(c rune) bool {
		return c == ':' || c == ';'
	}

	// split out each draw
	drawFieldsFunc := func(c rune) bool {
		return c == ',' || c == ' '
	}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		game := new(gameData)

		mainFields := strings.FieldsFunc(line, mainFieldsFunc)

		if game.id, err = strconv.Atoi(mainFields[0][5:]); err != nil {
			log.Fatal(err)
		}

		for _, d := range mainFields[1:] {
			drawFields := strings.FieldsFunc(d, drawFieldsFunc)
			draw := new(drawData)
			for idx, s := range drawFields {
				count := 0
				switch s {
				case "red":
					if count, err = strconv.Atoi(drawFields[idx-1]); err != nil {
						log.Fatal(err)
					}
					draw.red += count
				case "green":
					if count, err = strconv.Atoi(drawFields[idx-1]); err != nil {
						log.Fatal(err)
					}
					draw.green += count
				case "blue":
					if count, err = strconv.Atoi(drawFields[idx-1]); err != nil {
						log.Fatal(err)
					}
					draw.blue += count
				}
			}
			game.draws = append(game.draws, *draw)
		}
		games = append(games, *game)
	}
	return games
}

func part1() {
	var total = 0
	for _, game := range readGameData() {
		valid := true
		for _, draw := range game.draws {
			if draw.red > p1RedMax || draw.green > p1GreenMax || draw.blue > p1BlueMax {
				valid = false
				break
			}
		}
		if valid {
			total += game.id
		}
	}
	fmt.Printf("P1 Total: %d\n", total)
}
func part2() {
	var total = 0
	for _, game := range readGameData() {
		var (
			maxRed   = 0
			maxGreen = 0
			maxBlue  = 0
		)
		for _, draw := range game.draws {
			if draw.red > maxRed {
				maxRed = draw.red
			}
			if draw.green > maxGreen {
				maxGreen = draw.green
			}
			if draw.blue > maxBlue {
				maxBlue = draw.blue
			}
		}
		total += maxRed * maxGreen * maxBlue
	}
	fmt.Printf("P2 Total: %d\n", total)
}
func main() {
	part1()
	part2()
}
