package main

import (
	U "camel-cards/utils"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

const (
	CARD_INITIAL_VALUE = 2
	CARD_VALUES_AMOUNT = 13
)

type HandCombination uint8

const (
	HighCard HandCombination = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Hand  string
	Bid   int
	Power []int
}

func main() {
	file := filepath.Join("..", "assets", "cards.txt")
	hands, err := U.ReadFileLineByLine(file, ParseLineHand)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	OrderHands(hands)
	totalBidding := GetTotalBidding(hands)
	fmt.Printf("Total bidding: %v\n", totalBidding)
}

func ParseLineHand(line string) Hand {
	hand := strings.Split(line, " ")

	power := GetHandPower(hand[0])

	return Hand{
		hand[0],
		U.Atoi(hand[1]),
		power,
	}
}

func GetHandPower(hand string) (power []int) {
	duplicates := make([]int, CARD_VALUES_AMOUNT)

	for _, card := range hand {
		cardValue := GetCardValue(string(card))
		power = append(power, cardValue)

		duplicates[cardValue-CARD_INITIAL_VALUE]++
	}

	handCombination := GetHandCombination(duplicates)
	power = append([]int{int(handCombination)}, power...)

	return power
}

func GetCardValue(card string) int {
	switch card {
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		return U.Atoi(card)
	}
}

func GetHandCombination(duplicates []int) HandCombination {
	sort.Slice(duplicates, func(i, j int) bool {
		return duplicates[i] > duplicates[j]
	})

	switch {
	case duplicates[0] == 5:
		return FiveOfAKind
	case duplicates[0] == 4:
		return FourOfAKind
	case duplicates[0] == 3 && duplicates[1] == 2:
		return FullHouse
	case duplicates[0] == 3:
		return ThreeOfAKind
	case duplicates[0] == 2 && duplicates[1] == 2:
		return TwoPair
	case duplicates[0] == 2:
		return OnePair
	default:
		return HighCard
	}
}

func OrderHands(hands []Hand) []Hand {
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].Power[0] != hands[j].Power[0] {
			return hands[i].Power[0] < hands[j].Power[0]

		} else {
			for k := 1; k < len(hands[i].Power); k++ {
				if hands[i].Power[k] != hands[j].Power[k] {
					return hands[i].Power[k] < hands[j].Power[k]
				}
			}
		}

		return false
	})

	return hands
}

func GetTotalBidding(hands []Hand) int {
	total := 0

	for i, hand := range hands {
		total += hand.Bid * (i + 1)
	}

	return total
}
