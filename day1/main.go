package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	var total = 0

	for i := 0; reader.Scan(); i++ {
		line := reader.Text()
		var numString string

		// get first digit from the start
		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				numString += string(line[i])
				break
			}
		}

		// get first digit from the end
		for i := len(line) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				numString += string(line[i])
				break
			}
		}

		// check the string and sum total
		if num, err := strconv.Atoi(numString); err == nil {
			total += num
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("P1 Total: %d\n", total)
}

var numWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func isDigit(line string, idx int) (bool, string) {
	if _, err := strconv.Atoi(string(line[idx])); err == nil {
		return true, string(line[idx])
	}
	for word, digit := range numWords {
		if strings.HasPrefix(line[idx:], word) {
			return true, strconv.Itoa(digit)
		}
	}
	return false, ""
}
func part2() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	var total = 0

	for i := 0; reader.Scan(); i++ {
		line := reader.Text()
		var numString string

		// get first digit from the start
		for i := 0; i < len(line); i++ {
			if ok, str := isDigit(line, i); ok {
				numString += str
				break
			}
		}

		// get first digit from the end
		for i := len(line) - 1; i >= 0; i-- {
			if ok, str := isDigit(line, i); ok {
				numString += str
				break
			}
		}

		// check the string and sum total
		if num, err := strconv.Atoi(numString); err == nil {
			total += num
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("P2 Total: %d\n", total)
}

func main() {
	part1()
	part2()
}
