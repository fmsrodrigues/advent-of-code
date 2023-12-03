package main_test

import (
	"testing"

	GearRatios "gear-ratios"
)

func TestGearRatios(t *testing.T) {
	input := []string{"467..114..", "...*......", "..35..633.", "......#...", "617*......", ".....+.58.", "..592.....", "......755.", "...$.*....", ".664.598.."}

	numbers := []GearRatios.Number{
		{Value: 467, Positions: []GearRatios.Position{{0, 0}, {0, 1}, {0, 2}}},
		{Value: 114, Positions: []GearRatios.Position{{0, 5}, {0, 6}, {0, 7}}},
		{Value: 35, Positions: []GearRatios.Position{{2, 2}, {2, 3}}},
		{Value: 633, Positions: []GearRatios.Position{{2, 6}, {2, 7}, {2, 8}}},
		{Value: 617, Positions: []GearRatios.Position{{4, 0}, {4, 1}, {4, 2}}},
		{Value: 58, Positions: []GearRatios.Position{{5, 7}, {5, 8}}},
		{Value: 592, Positions: []GearRatios.Position{{6, 2}, {6, 3}, {6, 4}}},
		{Value: 755, Positions: []GearRatios.Position{{7, 6}, {7, 7}, {7, 8}}},
		{Value: 664, Positions: []GearRatios.Position{{9, 1}, {9, 2}, {9, 3}}},
		{Value: 598, Positions: []GearRatios.Position{{9, 5}, {9, 6}, {9, 7}}},
	}

	symbols := []GearRatios.Symbol{
		{42, []GearRatios.Position{{0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 4}, {2, 2}, {2, 3}, {2, 4}}},
		{35, []GearRatios.Position{{2, 5}, {2, 6}, {2, 7}, {3, 5}, {3, 7}, {4, 5}, {4, 6}, {4, 7}}},
		{42, []GearRatios.Position{{3, 2}, {3, 3}, {3, 4}, {4, 2}, {4, 4}, {5, 2}, {5, 3}, {5, 4}}},
		{43, []GearRatios.Position{{4, 4}, {4, 5}, {4, 6}, {5, 4}, {5, 6}, {6, 4}, {6, 5}, {6, 6}}},
		{36, []GearRatios.Position{{7, 2}, {7, 3}, {7, 4}, {8, 2}, {8, 4}, {9, 2}, {9, 3}, {9, 4}}},
		{42, []GearRatios.Position{{7, 4}, {7, 5}, {7, 6}, {8, 4}, {8, 6}, {9, 4}, {9, 5}, {9, 6}}},
	}

	symbolsNumbers := []GearRatios.SymbolNumbers{
		{
			Symbol: GearRatios.Symbol{42, []GearRatios.Position{{0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 4}, {2, 2}, {2, 3}, {2, 4}}},
			Numbers: []GearRatios.Number{
				{Value: 467, Positions: []GearRatios.Position{{0, 0}, {0, 1}, {0, 2}}},
				{Value: 35, Positions: []GearRatios.Position{{2, 2}, {2, 3}}},
			},
		},
		{
			Symbol: GearRatios.Symbol{35, []GearRatios.Position{{2, 5}, {2, 6}, {2, 7}, {3, 5}, {3, 7}, {4, 5}, {4, 6}, {4, 7}}},
			Numbers: []GearRatios.Number{
				{Value: 633, Positions: []GearRatios.Position{{2, 6}, {2, 7}, {2, 8}}},
			},
		},
		{
			Symbol: GearRatios.Symbol{42, []GearRatios.Position{{3, 2}, {3, 3}, {3, 4}, {4, 2}, {4, 4}, {5, 2}, {5, 3}, {5, 4}}},
			Numbers: []GearRatios.Number{
				{Value: 617, Positions: []GearRatios.Position{{4, 0}, {4, 1}, {4, 2}}},
			},
		},
		{
			Symbol: GearRatios.Symbol{43, []GearRatios.Position{{4, 4}, {4, 5}, {4, 6}, {5, 4}, {5, 6}, {6, 4}, {6, 5}, {6, 6}}},
			Numbers: []GearRatios.Number{
				{Value: 592, Positions: []GearRatios.Position{{6, 2}, {6, 3}, {6, 4}}},
			},
		},
		{
			Symbol: GearRatios.Symbol{36, []GearRatios.Position{{7, 2}, {7, 3}, {7, 4}, {8, 2}, {8, 4}, {9, 2}, {9, 3}, {9, 4}}},
			Numbers: []GearRatios.Number{
				{Value: 664, Positions: []GearRatios.Position{{9, 1}, {9, 2}, {9, 3}}},
			},
		},
		{
			Symbol: GearRatios.Symbol{42, []GearRatios.Position{{7, 4}, {7, 5}, {7, 6}, {8, 4}, {8, 6}, {9, 4}, {9, 5}, {9, 6}}},
			Numbers: []GearRatios.Number{
				{Value: 755, Positions: []GearRatios.Position{{7, 6}, {7, 7}, {7, 8}}},
				{Value: 598, Positions: []GearRatios.Position{{9, 5}, {9, 6}, {9, 7}}},
			},
		},
	}

	result := 4361

	t.Run("Get elements in line", func(t *testing.T) {
		got := struct {
			numbers []GearRatios.Number
			symbols []GearRatios.Symbol
		}{[]GearRatios.Number{}, []GearRatios.Symbol{}}

		for y, line := range input {
			numbersInLine, symbolsInLine := GearRatios.GetElementsInLine(line, y)

			got.numbers = append(got.numbers, numbersInLine...)
			got.symbols = append(got.symbols, symbolsInLine...)
		}

		if len(got.numbers) != len(numbers) {
			t.Errorf("Got %d numbers, want %d", len(got.numbers), len(numbers))
		}

		if len(got.symbols) != len(symbols) {
			t.Errorf("Got %d symbols, want %d", len(got.symbols), len(symbols))
		}

		assertNumbers(t, got.numbers, numbers)
		assertSymbols(t, got.symbols, symbols)
	})

	t.Run("Get valided numbers for symbols", func(t *testing.T) {
		got := GearRatios.GetValidedNumbersForSymbols(symbols, numbers)

		if len(got) != len(symbolsNumbers) {
			t.Errorf("Got %d symbols numbers, want %d", len(got), len(symbolsNumbers))
		}

		for i, symbolNumbers := range got {
			assertNumbers(t, symbolNumbers.Numbers, symbolsNumbers[i].Numbers)
			assertSymbols(t, []GearRatios.Symbol{symbolNumbers.Symbol}, []GearRatios.Symbol{symbolsNumbers[i].Symbol})
		}
	})

	t.Run("Get sum of all valid numbers", func(t *testing.T) {
		got := GearRatios.SumAllValidNumbers(symbolsNumbers)

		if got != result {
			t.Errorf("Got %d, want %d", got, result)
		}
	})
}

func assertNumbers(t testing.TB, got, want []GearRatios.Number) {
	t.Helper()

	for i, number := range got {
		if number.Value != want[i].Value {
			t.Errorf("Number: Got %d, want %d", number.Value, want[i].Value)
		}

		for j, position := range number.Positions {
			if position != want[i].Positions[j] {
				t.Errorf("Number: Got %v, want %v", position, want[i].Positions[j])
			}
		}
	}
}

func assertSymbols(t testing.TB, got, want []GearRatios.Symbol) {
	t.Helper()

	for i, symbol := range got {
		if symbol.Char != want[i].Char {
			t.Errorf("Symbol: Got %d, want %d", symbol.Char, want[i].Char)
		}

		for j, position := range symbol.AdjacentPositions {
			if position != want[i].AdjacentPositions[j] {
				t.Errorf("Symbol: Got %v, want %v", position, want[i].AdjacentPositions[j])
			}
		}
	}
}
