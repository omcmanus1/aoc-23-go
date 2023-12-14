package three

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// * search through each line in file
// * if number is adjacent to a symbol (excluding periods),
// * including horizontal/up/down/diagonal, add to total

func TaskOne() {
	file, err := os.Open("three/input.txt")
	if err != nil {
		fmt.Println("file open error:", err)
	}
	defer file.Close()

	symbolIndexes, numIndexes := populateRefs(`[^a-zA-Z0-9\s.]`, file)

	// loop through array of nums
	// check previous, current and following line for symbols
	// if symbol is covered by index range (plus neighbours), add to total

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

// only search for '*' in lines
// if '*' is adjacenent to exactly 2 numbers (on surrounding lines)
// multiply the two numbers and add to total

func TaskTwo() {
	file, err := os.Open("three/input.txt")
	if err != nil {
		fmt.Println("file open error:", err)
	}
	defer file.Close()

	symbolIndexes, numIndexes := populateRefs(`\*`, file)

	ratiosTotal := 0
	for gearInd, gearVals := range symbolIndexes {
		if len(symbolIndexes[gearInd]) == 0 {
			continue
		}
		for _, gearVal := range gearVals {
			// gearVal := symbolIndexes[gearInd][0]
			adjacentCount := []int{}
			for _, num := range numIndexes[gearInd] {
				if gearVal >= num["range"][0] && gearVal <= num["range"][1] {
					adjacentCount = append(adjacentCount, num["value"][0])
				}
			}
			if gearInd > 0 {
				for _, num := range numIndexes[gearInd-1] {
					if gearVal >= num["range"][0] && gearVal <= num["range"][1] {
						adjacentCount = append(adjacentCount, num["value"][0])
					}
				}
			}
			if gearInd < len(numIndexes)-1 {
				for _, num := range numIndexes[gearInd+1] {
					if gearVal >= num["range"][0] && gearVal <= num["range"][1] {
						adjacentCount = append(adjacentCount, num["value"][0])
					}
				}
			}
			if len(adjacentCount) == 2 {
				ratiosTotal += (adjacentCount[0] * adjacentCount[1])
			}
		}
	}

	fmt.Println("gear ratios total:", ratiosTotal)
}

func populateRefs(symbolPattern string, file *os.File) ([][]int, [][]map[string][]int) {
	numRegex := regexp.MustCompile(`(\d+)`)
	symbolRegex := regexp.MustCompile(`[^a-zA-Z0-9\s.]`)

	// indexes of symbol occurrences
	// one sub-slice per line with all symbol occurrences (multiple or none included)
	symbolIndexes := [][]int{}

	// slice of number occurrences
	// containing slices of maps (one sub-slice per line, one map per number)
	// map contains value (single item slice). plus indices [start, end]
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
	return symbolIndexes, numIndexes
}
