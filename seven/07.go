package seven

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/omcmanus1/aoc-23/utils"
)

// Order list of hands (similar to poker) in strength
// A = high, 2 = low
// If hands match, rank by the highest first non-matching card
// Multiply bid by rank (lowest rank = 1), add all together

type Hand struct {
	cards string
	bid   int
	score int
}

const cardLabels = "23456789TJQKA"
const newCardLabels = "J23456789TQKA"

func TaskOne() {
	scanner, file := utils.GetFileScanner("seven/test.txt")
	defer file.Close()

	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, " ")
		cards := lineSlice[0]
		bid := utils.StringToInt(lineSlice[1])
		hands = append(hands, Hand{cards, bid, 0})
	}
	hands = calculateScore(hands)
	hands = sortHands(hands)
	winnings := 0
	for i, hand := range hands {
		fmt.Printf("cards: %v || score: %v * %v = %v\n", hand, hand.bid, i+1, hand.bid*(i+1))
		winnings += hand.bid * (i + 1)
	}

	fmt.Println(winnings)
}

func TaskTwo() {
	scanner, file := utils.GetFileScanner("seven/test.txt")
	defer file.Close()

	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, " ")
		cards := lineSlice[0]
		bid := utils.StringToInt(lineSlice[1])
		hands = append(hands, Hand{cards, bid, 0})
	}
	hands = calculateNewScore(hands)
	hands = sortHands(hands)
	winnings := 0
	for i, hand := range hands {
		fmt.Printf("cards: %v || score: %v * %v = %v\n", hand, hand.bid, i+1, hand.bid*(i+1))
		winnings += hand.bid * (i + 1)
	}

	fmt.Println(winnings)
}

func calculateScore(hands []Hand) []Hand {
	for i := 0; i < len(hands); i++ {
		threeKind := false
		onePair := false
		for _, label := range cardLabels {
			labelStr := string(label)
			repeatCount := strings.Count(hands[i].cards, labelStr)
			rankScores(hands, i, &repeatCount, &onePair, &threeKind)
		}
	}
	return hands
}

func calculateNewScore(hands []Hand) []Hand {
	for i := 0; i < len(hands); i++ {
		threeKind := false
		onePair := false
		jokerCount := strings.Count(hands[i].cards, "J")
		fmt.Println(jokerCount)
		for _, label := range newCardLabels {
			labelStr := string(label)
			repeatCount := strings.Count(hands[i].cards, labelStr) + jokerCount
			rankScores(hands, i, &repeatCount, &onePair, &threeKind)
			jokerCount -= 1
		}
	}
	return hands
}

func rankScores(hands []Hand, i int, repeatCount *int, onePair, threeKind *bool) {
	switch *repeatCount {
	case 5:
		hands[i].score = 6
	case 4:
		hands[i].score = 5
	case 3:
		if *onePair {
			hands[i].score = 4
		} else {
			hands[i].score = 3
		}
		*threeKind = true
	case 2:
		if *threeKind {
			hands[i].score = 4
		} else if *onePair {
			hands[i].score = 2
		} else {
			hands[i].score = 1
			*onePair = true
		}
	}
}

func sortHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score == hands[j].score {
			for k := 0; k < len(hands[i].cards); k++ {
				index1 := strings.Index(cardLabels, string(hands[i].cards[k]))
				index2 := strings.Index(cardLabels, string(hands[j].cards[k]))
				if index1 != index2 {
					return index1 > index2
				}
			}
			return false
		}
		return hands[i].score > hands[j].score
	})
	slices.Reverse(hands)
	return hands
}
