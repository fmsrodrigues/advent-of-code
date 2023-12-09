package main

import (
	"fmt"
	"log"
	U "mirage-maintenance/utils"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	file := filepath.Join("..", "assets", "map.txt")
	numbers, err := U.ReadFileLineByLine(file, ExtractNumbersLines)
	if err != nil {
		log.Fatal(err)
	}

	reports := []bool{false, true}
	for _, report := range reports {
		fmt.Printf("Reverse report: %t, Extrapolated sum: %d\n", report, HistoryReport(numbers, report))
	}
}

func HistoryReport(numbers [][]int, reverseNumber bool) (extrapolatedSum int) {
	reportNumbers := slices.Clone(numbers)

	var historyNumbers []int
	for _, number := range reportNumbers {
		if reverseNumber {
			slices.Reverse(number)
		}

		lastNumbers := GetLastNumbersSequences(number)
		nextNumber := CalculateNextNumber(lastNumbers)

		historyNumbers = append(historyNumbers, nextNumber)
	}

	for _, number := range historyNumbers {
		extrapolatedSum += number
	}

	return
}

func ExtractNumbersLines(s string) []int {
	line := strings.Split(s, " ")
	line = U.RemoveEmptyStrings(line)

	numbers := []int{}
	for _, l := range line {
		num := U.Atoi(l)
		numbers = append(numbers, num)
	}

	return numbers
}

func GetLastNumbersSequences(numbers []int) (lastNumberInRows []int) {
while:
	for {
		lastNumberInRows = append(lastNumberInRows, numbers[len(numbers)-1])

		numbers = generateNextNumberSequence(numbers)

		for _, number := range numbers {
			if number != 0 {
				continue while
			}
		}

		lastNumberInRows = append(lastNumberInRows, numbers[len(numbers)-1])
		break
	}

	return
}

func generateNextNumberSequence(numbers []int) (nextNumbersRow []int) {
	for i, number := range numbers {
		if i == 0 {
			continue
		}

		num := number - numbers[i-1]
		nextNumbersRow = append(nextNumbersRow, num)
	}

	return nextNumbersRow
}

func CalculateNextNumber(numbers []int) (nextNumber int) {
	slices.Reverse(numbers)

	for _, number := range numbers {
		nextNumber = number + nextNumber
	}

	return
}
