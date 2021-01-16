package main

import "testing"

func TestNumbersGame(t *testing.T) {

	tt := []struct {
		name    string
		numbers []int
		target  int
		found   bool
		steps   int
	}{
		{"one step", []int{2, 3, 1, 10, 4, 25}, 250, true, 1},
		{"two step", []int{2, 3, 1, 10, 4, 25}, 500, true, 2},
		{"three step", []int{2, 3, 1, 10, 4, 25}, 501, true, 3},
		{"four step", []int{2, 3, 1, 10, 4, 25}, 860, true, 4},
		{"five step", []int{2, 3, 1, 10, 4, 25}, 863, true, 5},
		{"target not reached", []int{100, 75, 50, 25, 10, 10}, 709, false, 5},
	}

	// t.Fatalf("this test failed and stopped running")
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			solution := NumbersGame(tc.numbers, tc.target)
			if solution.found != tc.found {
				t.Errorf("%v, expected %v solution for NumbersGame %v with target %v, returned %v", tc.name, tc.found, tc.numbers, tc.target, solution.found)
			}
			if solution.found && solution.answer != tc.target {
				t.Errorf("%v, expected %v answer for NumbersGame %v, returned %v", tc.name, tc.target, tc.numbers, solution.answer)
			}
			if solution.steps != tc.steps {
				t.Errorf("%v, expected solution in %v steps: %v took %v steps", tc.name, tc.steps, solution.eq, solution.steps)
			}
		})
	}
}
