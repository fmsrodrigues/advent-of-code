package main_test

import (
	"slices"
	"testing"

	M "mirage-maintenance"
)

func TestMirageMaincenance(t *testing.T) {
	input := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}

	t.Run("Run part 1", func(t *testing.T) {
		actual := runMain(input, false)
		expected := 114
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("Run part 2", func(t *testing.T) {
		actual := runMain(input, true)
		expected := 2
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

}

func runMain(numbers [][]int, reverseNumber bool) int {
	var historyNumbers []int
	for _, number := range numbers {
		if reverseNumber {
			slices.Reverse(number)
		}

		lastNumbers := M.GetLastNumbersSequences(number)
		nextNumber := M.CalculateNextNumber(lastNumbers)

		historyNumbers = append(historyNumbers, nextNumber)
	}

	var extrapolatedSum int
	for _, number := range historyNumbers {
		extrapolatedSum += number
	}

	return extrapolatedSum
}
