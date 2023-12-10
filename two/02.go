package two

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/omcmanus1/aoc-23/utils"
)

// determine if each game result is possible with amount of colors
// as defined by colorLimits map obj
// loop through each line and sum total of each color for all games

func TaskOne() {
	file, err := utils.GetFile("two/02-input.txt")
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

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

	pattern := `(\d+)\s+(red|green|blue)`
	regex := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(file)
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
			fmt.Println("possible game", "total:", total, "to add:", game)
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
	file, err := utils.GetFile("two/02-input.txt")
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	pattern := `(\d+)\s+(red|green|blue)`
	regex := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(file)
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
		// fmt.Println("game", game, "minimums:", colorMinimums, "power:", setPower)
		total += setPower
		game++
	}
	fmt.Println(total)
}
