package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var symbolRegex *regexp.Regexp
var schema schematic

const (
	maxRow = 140
	maxCol = 140
)

type schematic [maxRow][maxCol]string

func init() {
	symbolRegex = regexp.MustCompile("[^0-9.]")
	readSchematic()
}

func readSchematic() {
	var data []byte
	var err error

	if data, err = os.ReadFile("./input.txt"); err != nil {
		log.Fatal(err)
	}

	for ri, line := range strings.Split(string(data), "\n") {
		for ci, cha := range line {
			schema[ri][ci] = string(cha)
		}
	}
}

func completePartNumber(ri, ci int) int {
	// start pn string
	pnStr := schema[ri][ci]

	// check left
	for i := ci - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(schema[ri][i]); err != nil {
			break
		}
		pnStr = schema[ri][i] + pnStr
	}

	// check right
	for i := ci + 1; i < maxCol; i++ {
		if _, err := strconv.Atoi(schema[ri][i]); err != nil {
			break
		}
		pnStr = pnStr + schema[ri][i]
	}

	// parse and return
	if num, err := strconv.Atoi(pnStr); err == nil {
		return num
	}

	return 0
}

func findPartNumbers(ri, ci int) []int {
	var partNums []int
	var pn int

	if ci-1 >= 0 {
		if _, err := strconv.Atoi(schema[ri][ci-1]); err == nil {
			pn = completePartNumber(ri, ci-1)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	if ci+1 < maxCol {
		if _, err := strconv.Atoi(schema[ri][ci+1]); err == nil {
			pn = completePartNumber(ri, ci+1)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	if ri-1 >= 0 {
		if _, err := strconv.Atoi(schema[ri-1][ci]); err == nil {
			pn = completePartNumber(ri-1, ci)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	if ri+1 < maxRow {
		if _, err := strconv.Atoi(schema[ri+1][ci]); err == nil {
			pn = completePartNumber(ri+1, ci)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	if ri-1 >= 0 && ci-1 >= 0 {
		if _, err := strconv.Atoi(schema[ri-1][ci-1]); err == nil {
			pn = completePartNumber(ri-1, ci-1)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	if ri-1 >= 0 && ci+1 < maxCol {
		if _, err := strconv.Atoi(schema[ri-1][ci+1]); err == nil {
			pn = completePartNumber(ri-1, ci+1)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	if ri+1 < maxRow && ci-1 >= 0 {
		if _, err := strconv.Atoi(schema[ri+1][ci-1]); err == nil {
			pn = completePartNumber(ri+1, ci-1)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	if ri+1 < maxRow && ci+1 < maxCol {
		if _, err := strconv.Atoi(schema[ri+1][ci+1]); err == nil {
			pn = completePartNumber(ri+1, ci+1)
			if !slices.Contains(partNums, pn) {
				partNums = append(partNums, pn)
			}
		}
	}
	return partNums
}

func part1() {
	var total = 0
	for ri, row := range schema {
		for ci, str := range row {
			if symbolRegex.MatchString(str) {
				//fmt.Printf("Checking around symbol: %s (%d:%d)\n", str, ri, ci)
				for _, pn := range findPartNumbers(ri, ci) {
					//fmt.Printf("Adding PN: %d\n", pn)
					total += pn
				}
			}
		}
	}
	fmt.Printf("P1 Total: %d\n", total)
}

func part2() {
	var total = 0
	for ri, row := range schema {
		for ci, str := range row {
			if str == "*" {
				gears := findPartNumbers(ri, ci)
				// they're only gears if we have two
				if len(gears) == 2 {
					// compute ratio and add to total
					total += gears[0] * gears[1]
				}
			}
		}
	}
	fmt.Printf("P1 Total: %d\n", total)
}

func main() {
	part1()
	part2()
}
