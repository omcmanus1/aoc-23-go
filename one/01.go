package one

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/omcmanus1/aoc-23/utils"
)

/*
- Read through calibration doc to find calibration value on each line
- Combine first and last digit from each line as 2 digit number
- Return sum of all numbers
*/

func TaskOne() {
	file, err := utils.GetFile("one/01-input.txt")

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	sum := 0
	count := 0
	digits := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			strVal := string(char)
			if num, err := strconv.Atoi(strVal); err == nil {
				digits = append(digits, num)
			}
		}
		concatStr := strconv.Itoa(digits[0]) + strconv.Itoa(digits[len(digits)-1])
		concatNum, err := strconv.Atoi(concatStr)
		if err != nil {
			fmt.Println("unable to convert to int")
			return
		}
		sum += concatNum
		// fmt.Println("Line ", count, ":", line, "digits: ", digits, " = ", concatStr, "running total: ", sum)
		digits = []int{}
		count++
	}

	fmt.Println("Final Total:", sum)
}

/*
- Same as above, but include numbers spelled alphabetically (e.g. "nine")
- Add first and last in each line, return sum of all numbers
*/

func TaskTwo() {
	file, err := utils.GetFile("one/01-input.txt")

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	sum := 0
	count := 0
	digits := []map[int]int{}

	// reference object to search for strings
	intReference := map[string]int{
		"zero":  0,
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

	for scanner.Scan() {
		line := scanner.Text()

		// find starting index of substring in line (if exists)
		// find index of digit in line (if exists)
		// append to digits array in order of index

		for k, v := range intReference {
			firstInd := strings.Index(line, k)
			lastInd := strings.LastIndex(line, k)
			if firstInd > -1 {
				digits = append(digits, map[int]int{firstInd: v}, map[int]int{lastInd: v})
			}
			if lastInd > -1 {
				digits = append(digits, map[int]int{lastInd: v})
			}
		}

		for index, char := range line {
			strVal := string(char)
			if num, err := strconv.Atoi(strVal); err == nil {
				digits = append(digits, map[int]int{index: num})
			}
		}

		concatStr := strconv.Itoa(digits[0][1]) + strconv.Itoa(digits[len(digits)-1][1])
		_, err := strconv.Atoi(concatStr)
		if err != nil {
			fmt.Println("unable to convert to int")
			return
		}

		sort.Slice(digits, func(i, j int) bool {
			key1, key2 := utils.GetKey(digits[i]), utils.GetKey(digits[j])
			return key1 < key2
		})
		accStr := strconv.Itoa(utils.GetValue(digits[0])) + strconv.Itoa(utils.GetValue(digits[len(digits)-1]))
		accNum, err := strconv.Atoi(accStr)
		if err != nil {
			fmt.Println("unable to convert to int")
			return
		}

		sum += accNum
		// fmt.Println("Line", count, ":", line, digits, accNum)
		digits = []map[int]int{}
		count++
	}

	fmt.Println("Final Total:", sum)

}
