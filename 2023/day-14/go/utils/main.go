package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileLineByLine[T any](file string, fn func(string) T) ([]T, error) {
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

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	return num
}

func RemoveEmptyStrings(slice []string) []string {
	var result []string

	for _, v := range slice {
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}
