package main

import (
	"testing"

	"github.com/larschri/sudokusolver/solver"
)

func TestSolve(t *testing.T) {
	tests := map[string]string{
		"........1.....2.3....45.6....1.7...2.7.6.....58....7....3..9...4..5..8..85.7.....": "634987251795162438128453697941378562372615984586294713263849175417536829859721346",
		"1.......1.....2.3....45.6....1.7...2.7.6.....58....7....3..9...4..5..8..85.7.....": "no solution",
		"........1.......3....45.6....1.7...2.7.6.....58....7....3..9...4..5..8..85.7.....": "multiple solutions",
		".....8.7.45.9..........64.5.......2..6..8...32.8.976...9.34.........2.9.3.48...62": "123458976456973218789126435547631829961284753238597641692345187815762394374819562",
	}

	for board, expected := range tests {
		t.Run(board, func(t *testing.T) {
			got, err := solver.Solve(board)
			if err != nil {
				if err.Error() != expected {
					t.Error(err)
				}
				return
			}

			if got != expected {
				t.Error(got)
			}
		})
	}
}
