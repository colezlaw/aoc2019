package main

import "testing"

func TestCalcFuel(t *testing.T) {
	tt := []struct {
		title string
		mass  int
		fuel  int
	}{
		{"For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.", 12, 2},
		{"For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.", 14, 2},
		{"For a mass of 1969, the fuel required is 654.", 1969, 654},
		{"For a mass of 100756, the fuel required is 33583.", 100756, 33583},
	}

	for _, tc := range tt {
		t.Run(tc.title, func(t *testing.T) {
			actual := calcFuel(tc.mass)
			if actual != tc.fuel {
				t.Errorf("expected fuel of %d, got %d", tc.fuel, actual)
			}
		})
	}
}

func TestCalcTotalFuel(t *testing.T) {
	tt := []struct {
		title string
		mass  int
		fuel  int
	}{
		{"14 requires 2 fuel", 14, 2},
		{"1969 requires 966 fuel", 1969, 966},
		{"100756 requires 50346 fuel", 100756, 50346},
	}

	for _, tc := range tt {
		t.Run(tc.title, func(t *testing.T) {
			actual := calcTotalFuel(tc.mass)
			if actual != tc.fuel {
				t.Errorf("expected fuel of %d, got %d", tc.fuel, actual)
			}
		})
	}
}
