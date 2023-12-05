package main_test

import (
	"path/filepath"
	"testing"

	Seeds "seeds"
)

func TestSeeds(t *testing.T) {
	file := filepath.Join("..", "assets", "seeds_map_test.txt")

	seeds := []int{79, 14, 55, 13}
	maps := []Seeds.Map{
		{"seed", "soil", []Seeds.MapConvertion{{98, 50, 2}, {50, 52, 48}}},
		{"soil", "fertilizer", []Seeds.MapConvertion{{15, 0, 37}, {52, 37, 2}, {0, 39, 15}}},
		{"fertilizer", "water", []Seeds.MapConvertion{{53, 49, 8}, {11, 0, 42}, {0, 42, 7}, {7, 57, 4}}},
		{"water", "light", []Seeds.MapConvertion{{18, 88, 7}, {25, 18, 70}}},
		{"light", "temperature", []Seeds.MapConvertion{{77, 45, 23}, {45, 81, 19}, {64, 68, 13}}},
		{"temperature", "humidity", []Seeds.MapConvertion{{69, 0, 1}, {0, 1, 69}}},
		{"humidity", "location", []Seeds.MapConvertion{{56, 60, 37}, {93, 56, 4}}},
	}

	locations := []int{82, 43, 86, 35}

	t.Run("Parse input to Seeds and Maps", func(t *testing.T) {
		fileContent, err := Seeds.ReadFile(file, func(line string) string {
			return line
		})
		if err != nil {
			t.Fatalf("Error reading file: %v", err)
		}

		gotSeeds, gotMaps := Seeds.MapSeedsMap(fileContent)

		assertPrimitiveArray(t, gotSeeds, seeds)
		assertMaps(t, gotMaps, maps)
	})

	t.Run("Walk map convertions", func(t *testing.T) {
		gotLocations := Seeds.WalkMapConvertions(seeds, maps)

		assertPrimitiveArray(t, gotLocations, locations)
	})
}

func assertPrimitiveArray(t testing.TB, got, want []int) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %v seeds, want %v seeds", len(got), len(want))
	}

	for i, gotSeed := range got {
		if gotSeed != want[i] {
			t.Errorf("got %v, want %v", gotSeed, want[i])
		}
	}
}

func assertMaps(t testing.TB, got, want []Seeds.Map) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %v maps, want %v maps", len(got), len(want))
	}

	for i, gotMap := range got {
		if gotMap.Source != want[i].Source {
			t.Errorf("got %v, want %v", gotMap.Source, want[i].Source)
		}

		if gotMap.Destination != want[i].Destination {
			t.Errorf("got %v, want %v", gotMap.Destination, want[i].Destination)
		}

		for j, gotMapConvertion := range gotMap.Convertions {
			if gotMapConvertion.From != want[i].Convertions[j].From {
				t.Errorf("got %v, want %v", gotMapConvertion.From, want[i].Convertions[j].From)
			}

			if gotMapConvertion.To != want[i].Convertions[j].To {
				t.Errorf("got %v, want %v", gotMapConvertion.To, want[i].Convertions[j].To)
			}

			if gotMapConvertion.Converted != want[i].Convertions[j].Converted {
				t.Errorf("got %v, want %v", gotMapConvertion.Converted, want[i].Convertions[j].Converted)
			}
		}
	}
}
