package five

import (
	"bufio"
	"fmt"
	"regexp"

	"golang.org/x/exp/slices"

	"github.com/omcmanus1/aoc-23/utils"
)

// For the list of seeds, find the lowest corresponding location number
// Map structure:
// * [0]: destination number
// * [1]: source number
// * [2]: range (i.e. same conversion difference to source for x increments)
// If number not in map range, conversion difference == 1

// Aim: create slice of maps for each conversion
// * Search range start: <source>
// * Search range end: <source + range - 1>
// * Conversion difference
// If num in range, apply conversion - else return num

type MapObj map[string]int

func TaskOne() {
	scanner, file := utils.GetFileScanner("five/input.txt")
	defer file.Close()

	seeds, conversionRefs := populateRefs(scanner)
	outputArr := []int{}

	for _, s := range seeds {
		outputArr = append(outputArr, singlePass(s, conversionRefs))
	}

	lowestOutput := slices.Min(outputArr)
	fmt.Println(lowestOutput)
}

func singlePass(seed int, conversionRefs [][]MapObj) int {
	for i := range conversionRefs {
		for _, singleMap := range conversionRefs[i] {
			if seed <= singleMap["end"] && seed >= singleMap["start"] {
				seed += singleMap["conversion"]
				i++
				break
			}
			if i == len(conversionRefs)-1 {
				i++
			}
		}
	}
	return seed
}

func populateRefs(scanner *bufio.Scanner) ([]int, [][]MapObj) {
	var seeds []int
	seedsComplete := false
	var tempArr []MapObj
	refsByGroup := [][]MapObj{}

	numRegex := regexp.MustCompile(`(\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := numRegex.FindAllString(line, -1)

		if !seedsComplete {
			for item := range matches {
				seed := utils.StringToInt(matches[item])
				seeds = append(seeds, seed)
			}
			if line == "" {
				seedsComplete = true
			}
			continue
		}
		if len(matches) != 0 {
			first := utils.StringToInt(matches[0])
			second := utils.StringToInt(matches[1])
			third := utils.StringToInt(matches[2])
			lineMatch := MapObj{
				"start":      second,
				"end":        second + third - 1,
				"conversion": first - second,
			}
			tempArr = append(tempArr, lineMatch)
		}
		if line == "" && len(tempArr) != 0 {
			refsByGroup = append(refsByGroup, tempArr)
			tempArr = []MapObj{}
		}
	}

	if len(tempArr) != 0 {
		refsByGroup = append(refsByGroup, tempArr)
	}

	return seeds, refsByGroup
}
