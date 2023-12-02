package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var nonNumericalRegex = regexp.MustCompile(`[^0-9 ]+`)

var numbers = map[string]string{
	"zero":  "z0o",
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	file := []string{"documents", "calibration.txt"}
	path := filepath.Join(file...)

	calibrationValue := GetCalibrationResult(path, GetCalibrationValue)
	fmt.Printf("Calibration value: %d\n", calibrationValue)

	calibrationMethodWithNumericalStrings := func(s string) int {
		s = ReplaceNumericalString(s)
		return GetCalibrationValue(s)
	}
	calibrationValueWithNumericalStrings := GetCalibrationResult(path, calibrationMethodWithNumericalStrings)
	fmt.Printf("Calibration value with numerical strings: %d\n", calibrationValueWithNumericalStrings)
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

func ReplaceNumericalString(s string) string {
	for k, v := range numbers {
		s = strings.ReplaceAll(s, k, v)
	}

	return s
}

func GetCalibrationValue(s string) int {
	s = nonNumericalRegex.ReplaceAllString(s, "")

	rawValue := fmt.Sprintf("%s%s", s[0:1], s[len(s)-1:])

	if num, err := strconv.Atoi(rawValue); err == nil {
		return num
	}

	log.Fatalf("Failed to convert %s to int", rawValue)
	return 0
}

func SumDigits(d []int) int {
	sum := 0

	for _, v := range d {
		sum += v
	}

	return sum
}

func GetCalibrationResult(s string, calibrationMethod func(string) int) int {
	calibrationValues, err := ReadFile[int](s, calibrationMethod)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	return SumDigits(calibrationValues)
}
