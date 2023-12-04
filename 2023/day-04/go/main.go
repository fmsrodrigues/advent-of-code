package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	scratchcards, err := ReadFile(filepath.Join("..", "assets", "scratchcards.txt"), GetNumbers)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	totalWorth := GetTotalWorthFromScratchcards(scratchcards)
	fmt.Printf("Total worth: %d\n", totalWorth)

	quantityOfScratchcardsInPile := GetQuantityOfScratchcardsInPile(scratchcards)
	fmt.Printf("Quantity of scratchcards in pile: %d\n", quantityOfScratchcardsInPile)
}

type Scratchcard struct {
	WinningNumbers []int
	OwnedNumbers   []int
	Worth          int
}

func (s *Scratchcard) getWorth() {
	var worth int

	for _, ownedNumber := range s.OwnedNumbers {
		for _, winningNumber := range s.WinningNumbers {
			if ownedNumber == winningNumber {
				if worth > 0 {
					worth *= 2
				} else {
					worth++
				}
			}
		}
	}

	s.Worth = worth
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

func GetNumbers(cardInput string) Scratchcard {
	card := strings.Split(cardInput, ":")
	numbers := strings.Split(card[1], "|")
	winnigNumbersInput := strings.Split(strings.TrimSpace(numbers[0]), " ")
	ownedNumbersInput := strings.Split(strings.TrimSpace(numbers[1]), " ")

	var winningNumbers []int
	for _, number := range winnigNumbersInput {
		if number == "" {
			continue
		}

		winningNumbers = append(winningNumbers, atoi(number))
	}

	var ownedNumbers []int
	for _, number := range ownedNumbersInput {
		if number == "" {
			continue
		}

		ownedNumbers = append(ownedNumbers, atoi(number))
	}

	scratchcard := Scratchcard{
		WinningNumbers: winningNumbers,
		OwnedNumbers:   ownedNumbers,
	}

	scratchcard.getWorth()
	return scratchcard
}

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	return num
}

func GetTotalWorthFromScratchcards(scratchcards []Scratchcard) int {
	var totalWorth int

	for _, scratchcard := range scratchcards {
		totalWorth += scratchcard.Worth
	}

	return totalWorth
}

func GetQuantityOfScratchcardsInPile(scratchcards []Scratchcard) int {
	quantityOfEachScratchcardInPile := []int{}

	nextCopy := 1
	for i, scratchcard := range scratchcards {
		if i >= len(quantityOfEachScratchcardInPile) {
			quantityOfEachScratchcardInPile = append(quantityOfEachScratchcardInPile, 0)
		}

		quantityOfEachScratchcardInPile[i] += 1

		for _, ownedNumber := range scratchcard.OwnedNumbers {
			for _, winningNumber := range scratchcard.WinningNumbers {

				if ownedNumber == winningNumber {
					copyPosition := i + nextCopy

					if copyPosition >= len(quantityOfEachScratchcardInPile)-1 {
						quantityOfEachScratchcardInPile = append(quantityOfEachScratchcardInPile, 0)
					}

					quantityOfEachScratchcardInPile[copyPosition] += 1 * quantityOfEachScratchcardInPile[i]
					nextCopy++

				}
			}
		}

		nextCopy = 1
	}

	var quantityOfScratchcardsInPile int

	for _, quantityOfEachScratchcard := range quantityOfEachScratchcardInPile {
		quantityOfScratchcardsInPile += quantityOfEachScratchcard
	}

	return quantityOfScratchcardsInPile
}
