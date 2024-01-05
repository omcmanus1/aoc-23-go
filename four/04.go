package four

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/omcmanus1/aoc-23/utils"
)

// Check each line of the file (scratchcard results)
// * Format: 5 winning nums | 8 nums to check for matches
// * 1 match = 1 point, then additional matches double the score
// * e.g. 16 points is the max score

func TaskOne() {
	scanner, file := utils.GetFile("four/input.txt")
	defer file.Close()

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		winningNums, testNums := getCardInputs(line)
		score := 0
		for _, num := range testNums {
			if slices.Contains(winningNums, num) && score == 0 {
				score = 1
			} else if slices.Contains(winningNums, num) {
				score *= 2
			}
		}
		fmt.Println("score:", score)
		totalScore += score
	}
	fmt.Println("Total Score:", totalScore)
}

// Instead of points, calculate:
// * no. of matches = copies of cards below winning card (1 of each)
// evaluate each card * num of copies you have

func TaskTwo() {
	scanner, file := utils.GetFile("four/test.txt")
	defer file.Close()

	type Score struct {
		original int
		copies   int
	}
	lineNum := 1
	copiesRef := map[int]Score{}

	for scanner.Scan() {
		line := scanner.Text()
		winningNums, testNums := getCardInputs(line)
		score := getCardScore(winningNums, testNums)
		copiesRef[lineNum] = Score{original: 1, copies: score}
		lineNum++
	}
	fmt.Println("Scores:", copiesRef)
	// for i :=
}

func getCardInputs(line string) ([]string, []string) {
	trimmed := strings.ReplaceAll(line, "  ", " ")
	_, after, _ := strings.Cut(trimmed, ": ")
	groups := strings.Split(after, " | ")
	winningNums := strings.Split(groups[0], " ")
	testNums := strings.Split(groups[1], " ")
	return winningNums, testNums
}

func getCardScore(winningNums, testNums []string) int {
	score := 0
	for _, num := range testNums {
		if slices.Contains(winningNums, num) {
			score++
		}
	}
	return score
}
