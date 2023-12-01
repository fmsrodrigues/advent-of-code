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
	file := []string{"documents", "calibration.txt"}
	calibrationValue := GetCalibrationValue(filepath.Join(file...))

	fmt.Printf("Calibration value: %d\n", calibrationValue)
}

func GetCalibrationValue(filename string) int {
	documentLines, err := ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var digits []int
	for _, line := range documentLines {
		digits = append(digits, GetDigits(line))
	}

	return SumDigits(digits)
}

func ReadFile(filename string) ([]string, error) {
	var lines []string

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type Digit struct {
	Value       int
	Initialized bool
}

func GetDigits(s string) int {
	first := Digit{0, false}
	last := Digit{0, false}

	characters := strings.Split(s, "")

	for _, c := range characters {
		if num, err := strconv.Atoi(c); err == nil {

			if !first.Initialized {
				first = Digit{num, true}
			}

			last = Digit{num, true}
		}
	}

	return (first.Value * 10) + last.Value
}

func SumDigits(d []int) int {
	sum := 0

	for _, v := range d {
		sum += v
	}

	return sum
}
