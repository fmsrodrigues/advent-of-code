package main_test

import (
	"path/filepath"
	"testing"
	Trebuchet "trebuchet"
)

type TrebuchetCalibrationTest struct {
	filepath string
	cases    []struct {
		in   string
		want int
	}
	calibrationResult           int
	withReplaceNumericalStrings bool
}

func TestTrebuchet(t *testing.T) {
	calibrationTestWithoutNumericalStrings := TrebuchetCalibrationTest{
		filepath.Join("documents", "calibration_test_01.txt"),
		[]struct {
			in   string
			want int
		}{
			{"1abc2", 12},
			{"pqr3stu8vwx", 38},
			{"a1b2c3d4e5f", 15},
			{"treb7uchet", 77},
		},
		142,
		false,
	}

	calibrationTestWithNumericalStrings := TrebuchetCalibrationTest{
		filepath.Join("documents", "calibration_test_02.txt"),
		[]struct {
			in   string
			want int
		}{
			{"two1nine", 29},
			{"eightwothree", 83},
			{"abcone2threexyz", 13},
			{"xtwone3four", 24},
			{"4nineeightseven2", 42},
			{"zoneight234", 14},
			{"7pqrstsixteen", 76},
		},
		281,
		true,
	}

	for _, calibrationTest := range []TrebuchetCalibrationTest{calibrationTestWithoutNumericalStrings, calibrationTestWithNumericalStrings} {
		calibrationMethod := func(s string) int {
			if calibrationTest.withReplaceNumericalStrings {
				s = Trebuchet.ReplaceNumericalString(s)
			}

			return Trebuchet.GetCalibrationValue(s)
		}

		t.Run("Read from file and return all the strings", func(t *testing.T) {
			got, err := Trebuchet.ReadFile[string](calibrationTest.filepath, func(s string) string { return s })
			if err != nil {
				t.Errorf("Failed to read file: %v", err)
			}

			for i, c := range calibrationTest.cases {
				if c.in != got[i] {
					t.Errorf("Trebuchet.ReadFile(%q, func(s string) string { return s }): got %q, want %q", calibrationTest.filepath, got, c.in)
				}
			}
		})

		for _, c := range calibrationTest.cases {
			t.Run("Get the first and last digit in a string and return the combined value", func(t *testing.T) {
				got := calibrationMethod(c.in)

				if got != c.want {
					t.Errorf("CalibrationMethod(%q): got %d, want %d", c.in, got, c.want)
				}
			})
		}

		t.Run("Sum the digits returned from the calibration method", func(t *testing.T) {
			digits := make([]int, len(calibrationTest.cases))

			for _, c := range calibrationTest.cases {
				digits = append(digits, c.want)
			}

			got := Trebuchet.SumDigits(digits)
			if got != calibrationTest.calibrationResult {
				t.Errorf("Trebuchet.SumDigits(%q): got %d, want %d", digits, got, calibrationTest.calibrationResult)
			}
		})

		t.Run("Get the calibration results from the document digits", func(t *testing.T) {
			got := Trebuchet.GetCalibrationResult(calibrationTest.filepath, calibrationMethod)
			if got != calibrationTest.calibrationResult {
				t.Errorf("Trebuchet.GetCalibrationResult(%q, %s): got %d, want %d", calibrationTest.filepath, "calibrationMethod", got, calibrationTest.calibrationResult)
			}
		})
	}
}
