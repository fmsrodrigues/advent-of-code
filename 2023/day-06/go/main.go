package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	U "wait-for-it/utils"
)

func main() {
	file := filepath.Join("..", "assets", "race.txt")
	race, err := U.ReadFileLineByLine(file, GetRaceConditions)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	winningScenarios := GetRaceWinningScenarios(race[0], race[1])
	totalWinningCombinations := GetTotalWinningCombinations(winningScenarios)
	fmt.Printf("Total winning combinations: %v\n", totalWinningCombinations)
}

func GetRaceConditions(line string) (conditions []int) {
	data := strings.Split(line, ":")
	data = strings.Split(strings.TrimSpace(data[1]), " ")
	data = U.RemoveEmptyStrings(data)

	for _, v := range data {
		conditions = append(conditions, U.Atoi(v))
	}

	return conditions
}

func GetRaceWinningScenarios(durations, distances []int) (scenarios []int) {
	for i, minDistance := range distances {
		scenarios = append(scenarios, 0)

		var distance int
		duration := durations[i]
		for min := 1; min < duration; min++ {
			max := duration - min
			distance = min * max

			if distance > minDistance {
				scenarios[i] = max - (min - 1)

				break
			}
		}

	}

	fmt.Printf("Scenarios: %v\n", scenarios)

	return scenarios
}

func GetTotalWinningCombinations(scenarios []int) (total int) {
	total = scenarios[0]

	for _, v := range scenarios[1:] {
		total *= v
	}

	return total
}
