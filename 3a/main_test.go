package main

import "testing"

func TestSamples(t *testing.T) {
	tt := []struct {
		title    string
		wire1    []string
		wire2    []string
		expected int
	}{
		{"verbose sample", []string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}, 6},
		{"sample 1", []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, 159},
		{"sample 2", []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}, 135},
	}

	for _, tc := range tt {
		t.Run(tc.title, func(t *testing.T) {
			actual := intersect(tc.wire1, tc.wire2)
			if actual != tc.expected {
				t.Errorf("expected distance of %d, got %d", tc.expected, actual)
			}
		})
	}
}
