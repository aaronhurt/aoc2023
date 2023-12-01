package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"strconv"
	"strings"
)

var digitMap = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}
var wordMap = map[string]int{
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

func isDigit(line string, idx int, m map[string]int) (bool, int) {
	for str, digit := range m {
		if strings.HasPrefix(line[idx:], str) {
			return true, digit
		}
	}
	return false, 0
}

func readLines() []string {
	var data []byte
	var err error

	if data, err = os.ReadFile("./input.txt"); err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

func part1() {
	var total = 0

	for _, line := range readLines() {
		var numString string
		var length = len(line)

		// get first digit from the start
		for i := 0; i < length; i++ {
			if ok, digit := isDigit(line, i, digitMap); ok {
				numString += strconv.Itoa(digit)
				break
			}
		}

		// get first digit from the end
		for i := length - 1; i >= 0; i-- {
			if ok, digit := isDigit(line, i, digitMap); ok {
				numString += strconv.Itoa(digit)
				break
			}
		}

		// check the concatenated string and sum total
		if num, err := strconv.Atoi(numString); err == nil {
			total += num
		} else {
			log.Fatal(err)
		}
	}

	fmt.Printf("P1 Total: %d\n", total)
}

func part2() {
	var total = 0
	var fullMap = maps.Clone(digitMap)
	maps.Copy(fullMap, wordMap)

	for _, line := range readLines() {
		var numString string
		var length = len(line)

		// get first digit from the start
		for i := 0; i < length; i++ {
			if ok, digit := isDigit(line, i, fullMap); ok {
				numString += strconv.Itoa(digit)
				break
			}
		}

		// get first digit from the end
		for i := length - 1; i >= 0; i-- {
			if ok, digit := isDigit(line, i, fullMap); ok {
				numString += strconv.Itoa(digit)
				break
			}
		}

		// check the concatenated string and sum total
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
