package main

import (
	"fmt"
	"log"
)

type solver struct {
	// contains ascii numbers (solved cells) and zero bytes (unsolved cells)
	board [81]byte
	// number of solved cells
	count int
	// bits[p][n] == true means that n is not a solution for position p
	bits [81][9]bool
}

func (s *solver) solve() error {
	for !s.solved() {
		if s.solveNumberForCell() != 0 {
			continue
		}
		if s.solveCellForNumber() != 0 {
			continue
		}
		return s.search()
	}
	return nil
}

func (s *solver) solved() bool {
	return s.count == 81
}

func (s *solver) setSolved(pos int, val byte) {
	if s.board[pos] != 0 {
		was := s.board[pos]
		s.board[pos] = 'X'
		log.Panicf("%d: %c -> %c: \n%s\n", pos, was, val, prettyBoard(s.board))
	}

	s.count++
	s.board[pos] = val

	for _, i := range rows.groups[rows.idxs[pos]] {
		s.bits[i][val-'1'] = true
	}
	for _, i := range cols.groups[cols.idxs[pos]] {
		s.bits[i][val-'1'] = true
	}
	for _, i := range boxes.groups[boxes.idxs[pos]] {
		s.bits[i][val-'1'] = true
	}
	for i := 0; i < 9; i++ {
		s.bits[pos][i] = true
	}
}

func (s *solver) solveCellForNumber() int {
	var res int
	for _, group := range groups {
		var alts [9][]int
		for _, p := range group {
			for n := 0; n < 9; n++ {
				if !s.bits[p][n] {
					alts[n] = append(alts[n], p)
				}
			}
		}
		for i, a := range alts {
			if len(a) == 1 {
				s.setSolved(a[0], byte('1'+i))
				res++
				// The loop above might yield multiple solutions for the
				// same cell if the board state is illegal.
				// Break here so we only process the first solution.
				break
			}
		}
	}
	return res
}

func (s *solver) solveNumberForCell() int {
	var res int
outer:
	for i, c := range s.board {
		if c != 0 {
			continue //already solved
		}

		var a byte
		for j, b := range s.bits[i] {
			if b {
				continue // already eliminated
			}
			if a != 0 {
				continue outer // multiple alternatives
			}
			a = byte('1' + j)
		}

		if a != 0 {
			s.setSolved(i, a)
			res++
		}
	}
	return res
}

func (s *solver) search() error {
	minAlts := 10
	var cell int
	for i, c := range s.board {
		if c != 0 {
			continue //already solved
		}

		var alts int
		for _, b := range s.bits[i] {
			if !b {
				alts++
			}
		}

		if alts == 0 {
			return fmt.Errorf("no solution")
		}

		if alts < minAlts {
			minAlts = alts
			cell = i

			if alts == 2 {
				break
			}
		}
	}

	var solution solver

	for n, b := range s.bits[cell] {
		if b {
			continue
		}

		var attempt solver = *s
		attempt.setSolved(cell, byte('1'+n))
		if err := attempt.solve(); err != nil {
			continue
		}

		// TODO: possible optimisation; disable check for muliple
		// solutions and return here

		if solution.solved() {
			return fmt.Errorf("multiple solutions")
		}
		solution = attempt
	}

	if !solution.solved() {
		return fmt.Errorf("no solution")
	}

	*s = solution
	return nil
}
