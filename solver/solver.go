package solver

import (
	"fmt"
	"log"
)

type solver struct {
	// contains ascii numbers (solved cells) and zero bytes (unsolved cells)
	board [81]byte
	// number of solved cells
	solvedCells int
	// eliminated[p][n] == true means that n is not a solution for position p
	eliminated [81][9]bool
}

func Solve(s string) (string, error) {
	if len(s) != 81 {
		return "", fmt.Errorf("illegal line length: %d", len(s))
	}

	var sl solver
	for i, c := range s {
		if c != '.' {
			sl.setSolved(i, byte(c))
		}
	}

	return sl.solve()
}

func (s *solver) solve() (string, error) {
	for !s.solved() {
		if s.solveNumberForCell() != 0 {
			continue
		}
		if s.solveCellForNumber() != 0 {
			continue
		}
		return s.search()
	}
	return string(s.board[:]), nil
}

func (s *solver) solved() bool {
	return s.solvedCells == 81
}

func (s *solver) setSolved(pos int, val byte) {
	if s.board[pos] != 0 {
		was := s.board[pos]
		s.board[pos] = 'X'
		log.Panicf("%d: %c -> %c: \n%s\n", pos, was, val, s)
	}

	s.solvedCells++
	s.board[pos] = val

	for _, i := range neighbours[pos] {
		s.eliminated[i][val-'1'] = true
	}

	for i := 0; i < 9; i++ {
		s.eliminated[pos][i] = true
	}
}

func (s *solver) solveCellForNumber() int {
	var res int
	for _, group := range groups {
		var alts [9][]int
		for _, p := range group {
			for n := 0; n < 9; n++ {
				if !s.eliminated[p][n] {
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
		for j, b := range s.eliminated[i] {
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

func (s *solver) search() (string, error) {
	minAlts := 10
	var cell int
	for i, c := range s.board {
		if c != 0 {
			continue //already solved
		}

		var alts int
		for _, b := range s.eliminated[i] {
			if !b {
				alts++
			}
		}

		if alts == 0 {
			return "", fmt.Errorf("no solution")
		}

		if alts < minAlts {
			minAlts = alts
			cell = i

			if alts == 2 {
				break
			}
		}
	}

	var solution string

	for n, b := range s.eliminated[cell] {
		if b {
			continue
		}

		var attempt solver = *s
		attempt.setSolved(cell, byte('1'+n))
		s, err := attempt.solve()
		if err != nil {
			continue
		}

		// TODO: possible optimisation; disable check for muliple
		// solutions and return here

		if solution != "" {
			return "", fmt.Errorf("multiple solutions")
		}
		solution = s
	}

	if solution == "" {
		return "", fmt.Errorf("no solution")
	}

	return solution, nil
}

func (s *solver) String() string {
	board := []byte(`c c c  c c c  c c c
c c c  c c c  c c c
c c c  c c c  c c c

c c c  c c c  c c c
c c c  c c c  c c c
c c c  c c c  c c c

c c c  c c c  c c c
c c c  c c c  c c c
c c c  c c c  c c c
`)
	var j int
	for i, c := range board {
		if c == 'c' {
			board[i] = s.board[j]
			j++
		}
	}

	return string(board)
}
