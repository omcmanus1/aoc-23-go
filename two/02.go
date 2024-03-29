package two

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/omcmanus1/aoc-23/utils"
)

// determine if each game result is possible with amount of colors
// as defined by colorLimits map obj
// loop through each line and sum total of each color for all games

func TaskOne() {
	scanner, file := utils.GetFileScanner("two/input.txt")
	defer file.Close()

	colorLimits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	colorCounts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	regex := regexp.MustCompile(`(\d+)\s+(red|green|blue)`)

	game := 1
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		groups := strings.Split(line, ";")
		toAdd := []bool{}
		for _, group := range groups {
			matches := regex.FindAllStringSubmatch(group, -1)
			for _, match := range matches {
				color := match[2]
				if number, err := strconv.Atoi(match[1]); err == nil {
					colorCounts[color] = number
				}
			}
			if colorCounts["red"] <= colorLimits["red"] && colorCounts["green"] <= colorLimits["green"] && colorCounts["blue"] <= colorLimits["blue"] {
				toAdd = append(toAdd, true)
			} else {
				toAdd = append(toAdd, false)
			}
		}
		if !slices.Contains(toAdd, false) {
			total += game
		}
		game++
	}
	fmt.Println(total)
}

/*
 * Determine minimum number of cubes for each color for a possible game
 * Multiply red * blue * green min numbers to get the power of each set
 * Sum each set together for the answer
 */

func TaskTwo() {
	scanner, file := utils.GetFileScanner("two/input.txt")
	defer file.Close()

	regex := regexp.MustCompile(`(\d+)\s+(red|green|blue)`)

	game := 1
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		groups := strings.Split(line, ";")
		colorMinimums := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, group := range groups {
			matches := regex.FindAllStringSubmatch(group, -1)
			for _, match := range matches {
				color := match[2]
				if number, err := strconv.Atoi(match[1]); err == nil {
					if number > colorMinimums[color] {
						colorMinimums[color] = number
					}
				}
			}
		}
		setPower := colorMinimums["red"] * colorMinimums["green"] * colorMinimums["blue"]
		total += setPower
		game++
	}
	fmt.Println(total)
}
