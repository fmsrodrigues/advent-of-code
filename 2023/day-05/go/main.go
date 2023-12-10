package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type SeedsRangeMaps struct {
	SeedsRange []int
	Maps       []Map
}

type Map struct {
	Source      string
	Destination string
	Convertions []MapConvertion
}

type MapConvertion struct {
	From      int
	To        int
	Converted int
}

func main() {
	file := filepath.Join("..", "assets", "seeds_map.txt")
	lines, err := ReadFile(file, func(line string) string {
		return line
	})
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	Part1(lines)
}

func Part1(lines []string) {
	seeds, maps := MapSeedsMap(lines)

	locations := WalkMapConvertions(seeds, maps)
	fmt.Printf("Locations: %v\n", locations)

	fmt.Printf("Lowest location: %v\n", slices.Min(locations))
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

func MapSeedsMap(seedsMaps []string) ([]int, []Map) {
	var seeds []int
	var maps []Map

	var currentMap Map
	for _, seedMap := range seedsMaps {
		if strings.TrimSpace(seedMap) == "" {
			if currentMap.Source != "" {
				maps = append(maps, currentMap)
				currentMap = Map{}
			}

			continue

		} else if strings.Contains(seedMap, "seeds:") {
			seedsInfo := strings.Split(seedMap, ":")
			seedsInfo = strings.Split(strings.TrimSpace(seedsInfo[1]), " ")

			for _, seed := range seedsInfo {
				seeds = append(seeds, atoi(seed))
			}

		} else if strings.Contains(seedMap, "map:") {
			mapInfo := strings.Split(seedMap, " ")
			mapInfo = strings.Split(strings.TrimSpace(mapInfo[0]), "-to-")

			currentMap = Map{
				Source:      mapInfo[0],
				Destination: mapInfo[1],
				Convertions: []MapConvertion{},
			}

		} else {
			convertionInfo := strings.Split(seedMap, " ")

			currentMap.Convertions = append(currentMap.Convertions, MapConvertion{
				From:      atoi(convertionInfo[1]),
				To:        atoi(convertionInfo[0]),
				Converted: atoi(convertionInfo[2]),
			})
		}
	}

	if currentMap.Source != "" {
		maps = append(maps, currentMap)
	}

	return seeds, maps
}

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	return num
}

func WalkMapConvertions(seeds []int, maps []Map) (locations []int) {
	for _, seed := range seeds {
		currentValue := seed

		for _, currentMap := range maps {
			for _, convertion := range currentMap.Convertions {
				if currentValue >= convertion.From && currentValue <= convertion.From+convertion.Converted {
					currentValue = currentValue + (convertion.To - convertion.From)
					break
				}
			}
		}

		locations = append(locations, currentValue)
	}

	return locations
}
