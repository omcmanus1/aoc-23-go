package five

import (
	"bufio"
	"fmt"
	"regexp"

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
	scanner, file := utils.GetFileScanner("five/test.txt")
	defer file.Close()

	seeds, conversionRefs := populateRefs(scanner)
	fmt.Printf("%+v\n", seeds)
	fmt.Printf("%+v\n", conversionRefs)
	fmt.Printf("%+v\n", len(conversionRefs))

	// for _, s := range seeds {
	// 	singlePass(s, conversionRefs)
	// }

}

// func singlePass(seed int, conversionRefs [][]MapObj) int {
// 	fmt.Println(len(conversionRefs))
// 	// for i, group := range conversionRefs {
// 	// 	fmt.Println(i, group)
// 	// }
// 	return 0
// }

func populateRefs(scanner *bufio.Scanner) ([]int, [][]MapObj) {
	var seeds []int
	var tempArr []MapObj
	refsByGroup := [][]MapObj{}

	numRegex := regexp.MustCompile(`(\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := numRegex.FindAllString(line, -1)

		if len(seeds) == 0 {
			for item := range matches {
				seed := utils.StringToInt(matches[item])
				seeds = append(seeds, seed)
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
