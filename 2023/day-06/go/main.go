package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	U "wait-for-it/utils"
)

func main() {
	races()
	race()
}

func races() {
	file := filepath.Join("..", "assets", "race.txt")
	race, err := U.ReadFileLineByLine(file, GetRacesConditions)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	winningScenarios := GetRaceWinningScenarios(race[0], race[1])
	totalWinningCombinations := GetTotalWinningCombinations(winningScenarios)
	fmt.Printf("Total winning combinations: %v\n", totalWinningCombinations)
}

func race() {
	file := filepath.Join("..", "assets", "race.txt")
	newRace, err := U.ReadFileLineByLine(file, GetRaceCondition)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	winningScenarios := GetRaceWinningScenarios([]int{newRace[0]}, []int{newRace[1]})
	fmt.Printf("Total winning scenarios: %v\n", winningScenarios[0])
}

func GetRacesConditions(line string) (conditions []int) {
	data := strings.Split(line, ":")
	data = strings.Split(strings.TrimSpace(data[1]), " ")
	data = U.RemoveEmptyStrings(data)

	for _, v := range data {
		conditions = append(conditions, U.Atoi(v))
	}

	return conditions
}

func GetRaceCondition(line string) (condition int) {
	data := strings.Split(line, ":")
	raceCondition := strings.ReplaceAll(data[1], " ", "")

	return U.Atoi(raceCondition)
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

	return scenarios
}

func GetTotalWinningCombinations(scenarios []int) (total int) {
	total = scenarios[0]

	for _, v := range scenarios[1:] {
		total *= v
	}

	return total
}
