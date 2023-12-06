package main_test

import (
	"path/filepath"
	"reflect"
	"testing"

	W "wait-for-it"
	U "wait-for-it/utils"
)

func TestRaces(t *testing.T) {
	file := filepath.Join("..", "assets", "race_test.txt")
	duration := []int{7, 15, 30}
	distance := []int{9, 40, 200}
	winningScenarios := []int{4, 8, 9}
	totalWinningCombinations := 288

	t.Run("Get races duration and distance", func(t *testing.T) {
		got, err := U.ReadFileLineByLine(file, W.GetRacesConditions)
		if err != nil {
			t.Errorf("got %v want %v", got, err)
		}

		if !reflect.DeepEqual(duration, got[0]) {
			t.Errorf("got %v want %v", got[0], duration)
		}

		if !reflect.DeepEqual(distance, got[1]) {
			t.Errorf("got %v want %v", got[1], distance)
		}
	})

	t.Run("Get the amount of winning scenarios for each race", func(t *testing.T) {
		scenarios := W.GetRaceWinningScenarios(duration, distance)

		if !reflect.DeepEqual(winningScenarios, scenarios) {
			t.Errorf("got %v want %v", scenarios, winningScenarios)
		}

	})

	t.Run("Get the total winning combinations for all races", func(t *testing.T) {
		winningCombinations := W.GetTotalWinningCombinations(winningScenarios)

		if winningCombinations != totalWinningCombinations {
			t.Errorf("got %v want %v", winningCombinations, totalWinningCombinations)
		}
	})
}

func TestRace(t *testing.T) {
	file := filepath.Join("..", "assets", "race_test.txt")
	duration := 71530
	distance := 940200
	winningScenarios := 71503

	t.Run("Get races duration and distance", func(t *testing.T) {
		got, err := U.ReadFileLineByLine(file, W.GetRaceCondition)
		if err != nil {
			t.Errorf("got %v want %v", got, err)
		}

		if !reflect.DeepEqual(duration, got[0]) {
			t.Errorf("got %v want %v", got[0], duration)
		}

		if !reflect.DeepEqual(distance, got[1]) {
			t.Errorf("got %v want %v", got[1], distance)
		}
	})

	t.Run("Get the amount of winning scenarios", func(t *testing.T) {
		scenarios := W.GetRaceWinningScenarios([]int{duration}, []int{distance})

		if scenarios[0] != winningScenarios {
			t.Errorf("got %v want %v", scenarios, winningScenarios)
		}

	})
}
