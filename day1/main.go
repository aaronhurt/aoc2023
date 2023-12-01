package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func part2() {

}
func main() {
	part1()
	part2()
}
