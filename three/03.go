package three

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/omcmanus1/aoc-23/utils"
)

/*
* search through each line in file
* if number is adjacent to a symbol (not period),
* including horizontal/up/down/diagonal, add to total
 */

func TaskOne() {
	file, err := utils.GetFile("three/input.txt")
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	numRegex := regexp.MustCompile(`(\d+)`)
	symbolRegex := regexp.MustCompile(`[^a-zA-Z0-9\s.]`)

	// indexes of symbol occurrences
	// one sub-slice per line with all symbol occurrences (can be 1 or 0)
	symbolIndexes := [][]int{}

	// indexes of number occurrences
	// slice of slices of maps (one sub-slice per line, one map per number)
	// num value as key, slice of [start index, end index] as values
	numIndexes := [][]map[string][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		symbolMatches := symbolRegex.FindAllStringIndex(line, -1)
		tempSymbols := []int{}
		for _, v := range symbolMatches {
			tempSymbols = append(tempSymbols, v[0])
		}
		symbolIndexes = append(symbolIndexes, tempSymbols)

		numMatches := numRegex.FindAllStringIndex(line, -1)
		tempNums := []map[string][]int{}
		for _, match := range numMatches {
			matchValue, _ := strconv.Atoi(line[match[0]:match[1]])
			startIndex := 0
			if match[0] > 0 {
				startIndex = match[0] - 1
			}
			matchMap := map[string][]int{
				"value": {matchValue},
				"range": {startIndex, match[1]},
			}
			tempNums = append(tempNums, matchMap)
		}
		numIndexes = append(numIndexes, tempNums)
	}

	// loop through array of nums
	// check previous, current and following line for symbols
	// if symbol is covered by index range (plus start - 1), add to total

	partsTotal := 0
	for lineIndex := range numIndexes {
		for _, num := range numIndexes[lineIndex] {
			for _, symb := range symbolIndexes[lineIndex] {
				if symb >= num["range"][0] && symb <= num["range"][1] {
					partsTotal += num["value"][0]
				}
			}
			if lineIndex > 0 {
				for _, symb := range symbolIndexes[lineIndex-1] {
					if symb >= num["range"][0] && symb <= num["range"][1] {
						partsTotal += num["value"][0]
					}
				}
			}
			if lineIndex < len(symbolIndexes)-1 {
				for _, symb := range symbolIndexes[lineIndex+1] {
					if symb >= num["range"][0] && symb <= num["range"][1] {
						partsTotal += num["value"][0]
					}
				}
			}
		}
	}
	fmt.Println("parts total:", partsTotal)
}
