package main_test

import (
	CubeConundrum "cube-conundrum"
	"testing"
)

func TestCubeConundrum(t *testing.T) {
	// testFilePath := filepath.Join("..", "assets", "game_test.txt")

	bag := CubeConundrum.Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	cases := []struct {
		in         string
		game       CubeConundrum.Game
		isPossible bool
	}{
		{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			CubeConundrum.Game{
				Id: 1,
				GameSets: []CubeConundrum.GameSet{
					{Blue: 3, Red: 4},
					{Red: 1, Green: 2, Blue: 6},
					{Green: 2},
				},
			},
			true,
		},
		{
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			CubeConundrum.Game{
				Id: 2,
				GameSets: []CubeConundrum.GameSet{
					{Blue: 1, Green: 2},
					{Green: 3, Blue: 4, Red: 1},
					{Green: 1, Blue: 1},
				},
			},
			true,
		},
		{
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			CubeConundrum.Game{
				Id: 3,
				GameSets: []CubeConundrum.GameSet{
					{Green: 8, Blue: 6, Red: 20},
					{Blue: 5, Red: 4, Green: 13},
					{Green: 5, Red: 1},
				},
			},
			false,
		},
		{
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			CubeConundrum.Game{
				Id: 4,
				GameSets: []CubeConundrum.GameSet{
					{Green: 1, Red: 3, Blue: 6},
					{Green: 3, Red: 6},
					{Green: 3, Blue: 15, Red: 14},
				},
			},
			false,
		},
		{
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			CubeConundrum.Game{
				Id: 5,
				GameSets: []CubeConundrum.GameSet{
					{Red: 6, Blue: 1, Green: 3},
					{Blue: 2, Red: 1, Green: 2},
				},
			},
			true,
		},
	}

	sumIdGames := 8

	t.Run("Parse game sets", func(t *testing.T) {
		for _, c := range cases {
			got := CubeConundrum.ParseGameSets(c.in)

			if got.Id != c.game.Id {
				t.Errorf("ParseGameSets(%q): Id\n Got: %q\nWant: %q", c.in, got.Id, c.game.Id)
			}

			if len(got.GameSets) != len(c.game.GameSets) {
				t.Errorf("ParseGameSets(%q): Game sets length\n Got: %q\nWant: %q", c.in, got.GameSets, c.game.GameSets)
			}

			assertGameSets(t, got.GameSets, c.game.GameSets)
		}
	})

	t.Run("Check if game is possible", func(t *testing.T) {
		for _, c := range cases {
			got := CubeConundrum.CheckIfGameIsPossible(c.game, bag)

			if got != c.isPossible {
				t.Errorf("CheckIfGameIsPossible(%q): Is possible\n Got: %v\nWant: %v", c.game, got, c.isPossible)
			}
		}
	})

	t.Run("Sum id of all possible games", func(t *testing.T) {
		var games []CubeConundrum.Game
		for _, c := range cases {
			games = append(games, c.game)
		}

		got := CubeConundrum.SumIdOfAllPossibleGames(games, bag)

		if got != sumIdGames {
			t.Errorf("SumIdOfAllPossibleGames(%q): Sum of id\n Got: %v\nWant: %v", games, got, sumIdGames)
		}
	})
}

func assertGameSets(t testing.TB, got, want []CubeConundrum.GameSet) {
	t.Helper()

	for i, v := range want {
		if got[i].Blue != v.Blue || got[i].Green != v.Green || got[i].Red != v.Red {
			t.Errorf("GameSets doesn't match\n Got: %q\nWant: %q", got[i], v)
		}
	}
}
