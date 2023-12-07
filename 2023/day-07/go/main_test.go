package main_test

import (
	"reflect"
	"testing"

	C "camel-cards"
)

func TestCards(t *testing.T) {
	t.Run("Parse line hand to hand", func(t *testing.T) {
		input := []string{
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		}

		var got []C.Hand
		for _, line := range input {
			got = append(got, C.ParseLineHand(line))
		}

		want := []C.Hand{
			{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
			{"T55J5", 684, []int{3, 10, 5, 5, 11, 5}},
			{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
			{"KTJJT", 220, []int{2, 13, 10, 11, 11, 10}},
			{"QQQJA", 483, []int{3, 12, 12, 12, 11, 14}},
		}

		if !reflect.DeepEqual(got, want) {
			for i, hand := range got {
				t.Errorf("\nGot:  %v\nWant: %v", hand, want[i])
			}
		}
	})

	t.Run("Get card value", func(t *testing.T) {
		input := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
		want := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

		var got []int
		for _, card := range input {
			got = append(got, C.GetCardValue(card))
		}

		if !reflect.DeepEqual(got, want) {
			for i, value := range got {
				t.Errorf("\nGot:  %v\nWant: %v", value, want[i])
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
			got = append(got, C.GetHandCombination(duplicates))
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

	t.Run("Get card hand power", func(t *testing.T) {
		input := []string{
			"32T3K",
			"T55J5",
			"KK677",
			"KTJJT",
			"QQQJA",
		}

		var got [][]int
		for _, hand := range input {
			got = append(got, C.GetHandPower(hand))
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
		input := []C.Hand{
			{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
			{"T55J5", 684, []int{3, 10, 5, 5, 11, 5}},
			{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
			{"KTJJT", 220, []int{2, 13, 10, 11, 11, 10}},
			{"QQQJA", 483, []int{3, 12, 12, 12, 11, 14}},
		}

		got := C.OrderHands(input)

		want := []C.Hand{
			{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
			{"KTJJT", 220, []int{2, 13, 10, 11, 11, 10}},
			{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
			{"T55J5", 684, []int{3, 10, 5, 5, 11, 5}},
			{"QQQJA", 483, []int{3, 12, 12, 12, 11, 14}},
		}

		if !reflect.DeepEqual(got, want) {
			for i, hand := range got {
				t.Errorf("\nGot:  %v\nWant: %v", hand, want[i])
			}
		}
	})

	t.Run("Get game total bidding", func(t *testing.T) {
		input := []C.Hand{
			{"32T3K", 765, []int{1, 3, 2, 10, 3, 13}},
			{"KTJJT", 220, []int{2, 13, 10, 11, 11, 10}},
			{"KK677", 28, []int{2, 13, 13, 6, 7, 7}},
			{"T55J5", 684, []int{3, 10, 5, 5, 11, 5}},
			{"QQQJA", 483, []int{3, 12, 12, 12, 11, 14}},
		}

		got := C.GetTotalBidding(input)

		want := 6440

		if got != want {
			t.Errorf("\nGot:  %v\nWant: %v", got, want)
		}
	})
}
