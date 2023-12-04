package main_test

import (
	"testing"

	Scratchcards "scratchcards"
)

func TestScratchcards(t *testing.T) {
	input := []string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}

	want := []Scratchcards.Scratchcard{
		{WinningNumbers: []int{41, 48, 83, 86, 17}, OwnedNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}, Worth: 8},
		{WinningNumbers: []int{13, 32, 20, 16, 61}, OwnedNumbers: []int{61, 30, 68, 82, 17, 32, 24, 19}, Worth: 2},
		{WinningNumbers: []int{1, 21, 53, 59, 44}, OwnedNumbers: []int{69, 82, 63, 72, 16, 21, 14, 1}, Worth: 2},
		{WinningNumbers: []int{41, 92, 73, 84, 69}, OwnedNumbers: []int{59, 84, 76, 51, 58, 5, 54, 83}, Worth: 1},
		{WinningNumbers: []int{87, 83, 26, 28, 32}, OwnedNumbers: []int{88, 30, 70, 12, 93, 22, 82, 36}, Worth: 0},
		{WinningNumbers: []int{31, 18, 13, 56, 72}, OwnedNumbers: []int{74, 77, 10, 23, 35, 67, 36, 11}, Worth: 0},
	}

	totalWorth := 13

	quantityOfScratchardsInPile := 30

	t.Run("Get Scratchcards", func(t *testing.T) {
		for i, card := range input {
			got := Scratchcards.GetNumbers(card)

			for j, winningNumber := range got.WinningNumbers {
				if winningNumber != want[i].WinningNumbers[j] {
					t.Errorf("Card set %d | Winning number: got %d want %d", i+1, winningNumber, want[i].WinningNumbers[j])
				}
			}

			for j, ownedNumber := range got.OwnedNumbers {
				if ownedNumber != want[i].OwnedNumbers[j] {
					t.Errorf("Card set %d | Owned number: got %d want %d", i+1, ownedNumber, want[i].OwnedNumbers[j])
				}
			}

			if got.Worth != want[i].Worth {
				t.Errorf("Card set %d | Worth: got %d want %d", i+1, got.Worth, want[i].Worth)
			}
		}
	})

	t.Run("Get total worth from scratchcards", func(t *testing.T) {
		got := Scratchcards.GetTotalWorthFromScratchcards(want)

		if got != totalWorth {
			t.Errorf("Total worth: got %d want %d", got, totalWorth)
		}
	})

	t.Run("Get quantity of scratchcards in pile with copies", func(t *testing.T) {
		got := Scratchcards.GetQuantityOfScratchcardsInPile(want)

		if got != quantityOfScratchardsInPile {
			t.Errorf("Quantity of scratchcards in pile: got %d want %d", got, quantityOfScratchardsInPile)
		}
	})
}
