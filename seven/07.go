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
// Multiply bid by rank (lowest rank = 1)

type Hand struct {
	cards string
	bid   int
	score int
}

func TaskOne() {
	scanner, file := utils.GetFileScanner("seven/input.txt")
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

func calculateScore(hands []Hand) []Hand {
	for i := 0; i < len(hands); i++ {
		cardLabels := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
		threeKind := false
		onePair := false
		for _, label := range cardLabels {
			repeatCount := strings.Count(hands[i].cards, label)
			switch repeatCount {
			case 5:
				hands[i].score = 6
			case 4:
				hands[i].score = 5
			case 3:
				threeKind = true
				hands[i].score = 3
			case 2:
				if threeKind {
					hands[i].score = 4
				} else if onePair {
					hands[i].score = 2
				} else {
					hands[i].score = 1
					onePair = true
				}
			}
		}
	}
	return hands
}

func sortHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score == hands[j].score {
			return compareCards(hands[i].cards, hands[j].cards)
		}
		return hands[i].score > hands[j].score
	})
	slices.Reverse(hands)
	return hands
}

func compareCards(cards1, cards2 string) bool {
	cardLabels := "23456789TJQKA"
	for i := 0; i < len(cards1); i++ {
		index1 := strings.Index(cardLabels, string(cards1[i]))
		index2 := strings.Index(cardLabels, string(cards2[i]))
		if index1 != index2 {
			return index1 > index2
		}
	}
	return false
}
