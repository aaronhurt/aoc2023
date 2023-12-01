package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewScanner(file)

	var total = 0

	for i := 0; reader.Scan(); i++ {
		line := reader.Text()
		numString := ""

		// get first digit
		for _, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				numString += string(char)
				break
			}
		}
		// get last digit
		for _, char := range reverse(line) {
			if _, err := strconv.Atoi(string(char)); err == nil {
				numString += string(char)
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
	fmt.Printf("Total: %d\n", total)
}
