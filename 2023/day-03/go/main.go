package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

type Position struct {
	Y int
	X int
}

type Number struct {
	Value     int
	Positions []Position
}

type Symbol struct {
	Char              rune
	AdjacentPositions []Position
}

type SymbolNumbers struct {
	Symbol  Symbol
	Numbers []Number
}

func main() {
	file := filepath.Join("..", "assets", "gear_ratio.txt")

	lines, err := ReadFile(file, func(line string) string {
		return line
	})
	if err != nil {
		log.Fatalf("Could not read file: %q", err)
	}

	numbers := []Number{}
	symbols := []Symbol{}

	for y, line := range lines {
		lineNumbers, lineSymbols := GetElementsInLine(line, y)
		numbers = append(numbers, lineNumbers...)
		symbols = append(symbols, lineSymbols...)
	}

	symbolsNumbers := GetValidedNumbersForSymbols(symbols, numbers)

	sumWithoutGearRatioMultiplier := SumAllValidNumbers(symbolsNumbers, false)
	fmt.Printf("Sum of all valid numbers without gear ratio multiplier: %d\n", sumWithoutGearRatioMultiplier)

	sumWithGearRatioMultiplier := SumAllValidNumbers(symbolsNumbers, true)
	fmt.Printf("Sum of all valid numbers with gear ratio multiplier: %d\n", sumWithGearRatioMultiplier)
}

func ReadFile[T any](file string, fn func(string) T) ([]T, error) {
	var content []T

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content = append(content, fn(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

func GetElementsInLine(line string, y int) ([]Number, []Symbol) {
	var numbers []Number
	var symbols []Symbol

	digit := []rune{}
	positions := []Position{}

	// Add a dot at the end of the line to
	// handle the numbers in the last line position
	line = line + "."

	for x, char := range line {
		// Handle number
		if char >= '0' && char <= '9' {
			digit = append(digit, char)
			positions = append(positions, Position{y, x})

		} else {
			// Handle number ended
			if len(digit) > 0 {
				num, err := strconv.Atoi(string(digit))
				if err != nil {
					log.Fatalf("Could not convert digit:\nNumber: %q\nErrpr: %q", digit, err)
				}

				number := Number{num, positions}
				numbers = append(numbers, number)

				digit = []rune{}
				positions = []Position{}
			}

			// Handle symbol
			if char != '.' {
				symbol := Symbol{char, GetAdjacentPositions(Position{y, x})}
				symbols = append(symbols, symbol)
			}
		}
	}

	return numbers, symbols
}

func GetAdjacentPositions(position Position) []Position {
	adjacentPositions := []Position{
		{position.Y - 1, position.X - 1},
		{position.Y - 1, position.X},
		{position.Y - 1, position.X + 1},
		{position.Y, position.X - 1},
		{position.Y, position.X + 1},
		{position.Y + 1, position.X - 1},
		{position.Y + 1, position.X},
		{position.Y + 1, position.X + 1},
	}

	return adjacentPositions
}

func GetValidedNumbersForSymbols(symbols []Symbol, numbers []Number) []SymbolNumbers {
	symbolsNumbers := []SymbolNumbers{}

	length := 0

	for _, symbol := range symbols {
		symbolNumbers := SymbolNumbers{symbol, []Number{}}

		for _, number := range numbers {
			for _, numberPosition := range number.Positions {
				if slices.Contains(symbol.AdjacentPositions, numberPosition) {
					symbolNumbers.Numbers = append(symbolNumbers.Numbers, number)
					break
				}
			}
		}

		symbolsNumbers = append(symbolsNumbers, symbolNumbers)
		length += len(symbolNumbers.Numbers)
	}

	return symbolsNumbers
}

func SumAllValidNumbers(symbolsNumers []SymbolNumbers, useGearRatioMultiplier bool) int {
	sum := 0

	for _, symbolNumbers := range symbolsNumers {
		if useGearRatioMultiplier {
			if symbolNumbers.Symbol.Char == '*' && len(symbolNumbers.Numbers) == 2 {
				sum += symbolNumbers.Numbers[0].Value * symbolNumbers.Numbers[1].Value
			}
		} else {
			for _, number := range symbolNumbers.Numbers {
				sum += number.Value
			}
		}
	}

	return sum
}
