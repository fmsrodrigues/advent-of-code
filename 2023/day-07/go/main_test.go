package main_test

import (
	"reflect"
	"testing"

	C "camel-cards"
)

func TestGame(t *testing.T) {
	t.Run("Parse game lines", func(t *testing.T) {
		game := C.Game{}

		input := []string{
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		}

		var got []C.Hand
		for _, line := range input {
			got = append(got, game.ParseGameLine(line))
		}

		want := []C.Hand{
			{Hand: "32T3K", Bid: 765},
			{Hand: "T55J5", Bid: 684},
			{Hand: "KK677", Bid: 28},
			{Hand: "KTJJT", Bid: 220},
			{Hand: "QQQJA", Bid: 483},
		}

		if !reflect.DeepEqual(got, want) {
			for i, hand := range got {
				t.Errorf("\nGot:  %v\nWant: %v", hand, want[i])
			}
		}
	})

	t.Run("Get card hand combination", func(t *testing.T) {
		input := [][]int{
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 2, 1, 1},
			{0, 1, 2, 2, 0},
			{0, 0, 3, 1, 1},
			{0, 0, 2, 3, 0},
			{0, 1, 0, 0, 4},
			{0, 0, 0, 5, 0},
		}

		var got []C.HandCombination
		for _, duplicates := range input {
			got = append(got, C.Hand{}.GetHandCombination(duplicates))
		}

		want := []C.HandCombination{
			C.HighCard,
			C.HighCard,
			C.OnePair,
			C.TwoPair,
			C.ThreeOfAKind,
			C.FullHouse,
			C.FourOfAKind,
			C.FiveOfAKind,
		}

		if !reflect.DeepEqual(got, want) {
			for i, combination := range got {
				t.Errorf("\nGot:  %v\nWant: %v", combination, want[i])
			}
		}
	})
}

func TestBaseGame(t *testing.T) {
	t.Run("Get card value", func(t *testing.T) {
		game := C.Game{}

		input := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
		want := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

		var got []int
		for _, card := range input {
			got = append(got, game.GetCardValue(card))
		}

		if !reflect.DeepEqual(got, want) {
			for i, value := range got {
				t.Errorf("\nGot:  %v\nWant: %v", value, want[i])
			}
		}

	})

	t.Run("Get card hand power", func(t *testing.T) {
		input := []C.Hand{
			{Hand: "32T3K"},
			{Hand: "T55J5"},
			{Hand: "KK677"},
			{Hand: "KTJJT"},
			{Hand: "QQQJA"},
		}

		var got [][]int
		for _, hand := range input {
			got = append(got, hand.GetHandPower(&C.Game{}))
		}

		want := [][]int{
			{1, 3, 2, 10, 3, 13},
			{3, 10, 5, 5, 11, 5},
			{2, 13, 13, 6, 7, 7},
			{2, 13, 10, 11, 11, 10},
			{3, 12, 12, 12, 11, 14},
		}

		if !reflect.DeepEqual(got, want) {
			for i, power := range got {
				t.Errorf("\nGot:  %v\nWant: %v", power, want[i])
			}
		}
	})

	t.Run("Order all hands from weakest to strongest", func(t *testing.T) {
		game := C.Game{
			Hands: []C.Hand{
				{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
				{"T55J5", 684, []int{3, 10, 5, 5, 11, 5}},
				{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
				{"KTJJT", 220, []int{2, 13, 10, 11, 11, 10}},
				{"QQQJA", 483, []int{3, 12, 12, 12, 11, 14}},
			},
		}

		game.OrderHands()

		want := []C.Hand{
			{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
			{"KTJJT", 220, []int{2, 13, 10, 11, 11, 10}},
			{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
			{"T55J5", 684, []int{3, 10, 5, 5, 11, 5}},
			{"QQQJA", 483, []int{3, 12, 12, 12, 11, 14}},
		}

		if !reflect.DeepEqual(game.Hands, want) {
			for i, hand := range game.Hands {
				t.Errorf("\nGot:  %v\nWant: %v", hand, want[i])
			}
		}
	})

	t.Run("Get game total bidding", func(t *testing.T) {
		game := C.Game{
			Hands: []C.Hand{
				{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
				{"KTJJT", 220, []int{2, 13, 10, 11, 11, 10}},
				{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
				{"T55J5", 684, []int{3, 10, 5, 5, 11, 5}},
				{"QQQJA", 483, []int{3, 12, 12, 12, 11, 14}},
			},
		}

		game.GetTotalBidding()

		want := 6440

		if game.TotalBidding != want {
			t.Errorf("\nGot:  %v\nWant: %v", game.TotalBidding, want)
		}
	})

}

func TestJokerGame(t *testing.T) {
	t.Run("Get card value", func(t *testing.T) {
		game := C.Game{UseJokerRules: true}

		input := []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14}

		var got []int
		for _, card := range input {
			got = append(got, game.GetCardValue(card))
		}

		if !reflect.DeepEqual(got, want) {
			for i, value := range got {
				t.Errorf("\nGot:  %v\nWant: %v", value, want[i])
			}
		}

	})

	t.Run("Get card hand power", func(t *testing.T) {
		input := []C.Hand{
			{Hand: "32T3K"},
			{Hand: "T55J5"},
			{Hand: "KK677"},
			{Hand: "KTJJT"},
			{Hand: "QQQJA"},
		}

		var got [][]int
		for _, hand := range input {
			got = append(got, hand.GetHandPower(&C.Game{UseJokerRules: true}))
		}

		want := [][]int{
			{1, 3, 2, 10, 3, 13},
			{5, 10, 5, 5, 1, 5},
			{2, 13, 13, 6, 7, 7},
			{5, 13, 10, 1, 1, 10},
			{5, 12, 12, 12, 1, 14},
		}

		if !reflect.DeepEqual(got, want) {
			for i, power := range got {
				t.Errorf("\nGot:  %v\nWant: %v", power, want[i])
			}
		}
	})

	t.Run("Order all hands from weakest to strongest", func(t *testing.T) {
		game := C.Game{
			UseJokerRules: true,
			Hands: []C.Hand{
				{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
				{"T55J5", 684, []int{5, 10, 5, 5, 1, 5}},
				{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
				{"KTJJT", 220, []int{5, 13, 10, 1, 1, 10}},
				{"QQQJA", 483, []int{5, 12, 12, 12, 1, 14}},
			},
		}

		game.OrderHands()

		want := []C.Hand{
			{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
			{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
			{"T55J5", 684, []int{5, 10, 5, 5, 1, 5}},
			{"QQQJA", 483, []int{5, 12, 12, 12, 1, 14}},
			{"KTJJT", 220, []int{5, 13, 10, 1, 1, 10}},
		}

		if !reflect.DeepEqual(game.Hands, want) {
			for i, hand := range game.Hands {
				t.Errorf("\nGot:  %v\nWant: %v", hand, want[i])
			}
		}
	})

	t.Run("Get game total bidding", func(t *testing.T) {
		game := C.Game{
			UseJokerRules: true,
			Hands: []C.Hand{
				{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
				{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
				{"T55J5", 684, []int{5, 10, 5, 5, 1, 5}},
				{"QQQJA", 483, []int{5, 12, 12, 12, 1, 14}},
				{"KTJJT", 220, []int{5, 13, 10, 1, 1, 10}},
			},
		}

		game.GetTotalBidding()

		want := 5905

		if game.TotalBidding != want {
			t.Errorf("\nGot:  %v\nWant: %v", game.TotalBidding, want)
		}
	})
}
