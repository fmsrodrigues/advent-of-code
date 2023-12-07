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
	CARD_VALUES_AMOUNT = 15
)

type Game struct {
	Hands         []Hand
	TotalBidding  int
	UseJokerRules bool
}

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

func (g *Game) ReadGameFile() {
	file := filepath.Join("..", "assets", "cards.txt")
	hands, err := U.ReadFileLineByLine(file, g.ParseGameLine)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	g.Hands = hands
}

func (g Game) ParseGameLine(line string) Hand {
	hand := strings.Split(line, " ")

	return Hand{
		Hand: hand[0],
		Bid:  U.Atoi(hand[1]),
	}
}

func (g Game) GetCardValue(card string) int {
	switch card {
	case "T":
		return 10
	case "J":
		if g.UseJokerRules {
			return 1
		} else {
			return 11
		}
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

func (g *Game) CalculateHandsPower() {
	for i, hand := range g.Hands {
		g.Hands[i].Power = hand.GetHandPower(g)
	}
}

func (g *Game) OrderHands() {
	sort.SliceStable(g.Hands, func(i, j int) bool {
		if g.Hands[i].Power[0] != g.Hands[j].Power[0] {
			return g.Hands[i].Power[0] < g.Hands[j].Power[0]

		} else {
			for k := 1; k < len(g.Hands[i].Power); k++ {
				if g.Hands[i].Power[k] != g.Hands[j].Power[k] {
					return g.Hands[i].Power[k] < g.Hands[j].Power[k]
				}
			}
		}

		return false
	})
}

func (g *Game) GetTotalBidding() {
	for i, hand := range g.Hands {
		g.TotalBidding += hand.Bid * (i + 1)
	}
}

func (h Hand) GetHandCombination(duplicates []int) HandCombination {
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

func (h Hand) GetHandPower(game *Game) (power []int) {
	duplicates := make([]int, CARD_VALUES_AMOUNT)
	var jokers int

	for _, card := range h.Hand {
		cardValue := game.GetCardValue(string(card))
		power = append(power, cardValue)

		if game.UseJokerRules && cardValue == 1 {
			jokers++
		} else {
			duplicates[cardValue]++
		}
	}

	if jokers > 0 {
		highestDuplicate := 0
		highestDuplicateIndex := 0
		for i, value := range duplicates {
			if value > highestDuplicate {
				highestDuplicate = value
				highestDuplicateIndex = i
			}
		}

		duplicates[highestDuplicateIndex] += jokers
	}

	handCombination := h.GetHandCombination(duplicates)
	power = append([]int{int(handCombination)}, power...)

	return power
}

func main() {
	rules := []bool{false, true}
	for _, rule := range rules {
		game := Game{}
		game.ReadGameFile()
		game.UseJokerRules = rule

		game.CalculateHandsPower()
		game.OrderHands()

		game.GetTotalBidding()
		fmt.Printf("Using Joker rule: %v, Total bidding: %v\n", game.UseJokerRules, game.TotalBidding)
	}
}
