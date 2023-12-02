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

type Game struct {
	Id       int
	GameSets []GameSet
}

type GameSet struct {
	Red, Green, Blue int
}

type Bag GameSet

func main() {
	path := filepath.Join("..", "assets", "games.txt")
	games, err := ReadFile(path, ParseGameSets)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	bag := Bag{Red: 12, Green: 13, Blue: 14}
	sumIdGames := SumIdOfAllPossibleGames(games, bag)
	fmt.Printf("Sum of all possible games: %d\n", sumIdGames)
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

func ParseGameSets(gameLog string) Game {
	gameData := strings.Split(gameLog, ":")

	id, err := strconv.Atoi(strings.TrimPrefix(gameData[0], "Game "))
	if err != nil {
		log.Fatalf("Error converting game id to int:\nGame: %q\nError: %v", gameLog, err)
	}

	var gameSets []GameSet
	gameSetsText := strings.Split(gameData[1], ";")
	for i, gameSetText := range gameSetsText {
		cubes := strings.Split(gameSetText, ",")
		gameSets = append(gameSets, GameSet{})

		for _, cube := range cubes {
			cubeData := strings.Split(strings.TrimSpace(cube), " ")

			cubeQuantity, err := strconv.Atoi(cubeData[0])
			if err != nil {
				log.Fatalf("Error converting cube quantity to int:\nGame: %q\nCube: %q\nError: %v", gameLog, cubeData, err)
			}

			switch cubeData[1] {
			case "red":
				gameSets[i].Red = cubeQuantity
			case "green":
				gameSets[i].Green = cubeQuantity
			case "blue":
				gameSets[i].Blue = cubeQuantity
			default:
				log.Fatalf("Error parsing cube color:\nGame: %q\nCube: %q", gameLog, cube)
			}
		}
	}

	return Game{
		Id:       id,
		GameSets: gameSets,
	}
}

func CheckIfGameIsPossible(game Game, bag Bag) bool {
	for _, gameSet := range game.GameSets {
		if gameSet.Red > bag.Red || gameSet.Green > bag.Green || gameSet.Blue > bag.Blue {
			return false
		}
	}

	return true
}

func SumIdOfAllPossibleGames(games []Game, bag Bag) int {
	var sum int

	for _, game := range games {
		if CheckIfGameIsPossible(game, bag) {
			sum += game.Id
		}
	}

	return sum
}
