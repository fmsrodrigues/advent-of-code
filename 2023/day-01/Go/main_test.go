package main_test

import (
	"path/filepath"
	"testing"
	Trebuchet "trebuchet"
)

func TestTrebuchet(t *testing.T) {
	file := []string{"documents", "calibration_test.txt"}
	cases := []struct {
		in   string
		want int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}
	calibrationValue := 142

	t.Run("Read from file and return all the strings", func(t *testing.T) {
		filename := filepath.Join(file...)

		got, err := Trebuchet.ReadFile(filename)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		for i, c := range cases {
			if c.in != got[i] {
				t.Errorf("Trebuchet.ReadFile(%q): got %q, want %q", "calibration_test.txt", got, c.in)
			}
		}
	})

	for _, c := range cases {
		t.Run("Get the first and last digit in a string and return the combined value", func(t *testing.T) {
			got := Trebuchet.GetDigits(c.in)

			if got != c.want {
				t.Errorf("Trebuchet.GetDigits(%q): got %d, want %d", c.in, got, c.want)
			}
		})
	}

	t.Run("Sum the digits returned from GetDigits", func(t *testing.T) {
		digits := make([]int, len(cases))

		for _, c := range cases {
			digits = append(digits, c.want)
		}

		got := Trebuchet.SumDigits(digits)
		if got != calibrationValue {
			t.Errorf("Trebuchet.SumDigits(%q): got %d, want %d", digits, got, calibrationValue)
		}
	})

	t.Run("Get the calibration value from document", func(t *testing.T) {
		got := Trebuchet.GetCalibrationValue(filepath.Join(file...))
		if got != calibrationValue {
			t.Errorf("Trebuchet.GetCalibrationValue(%q): got %d, want %d", filepath.Join(file...), got, calibrationValue)
		}
	})

}
