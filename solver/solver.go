package solver

import (
	"fmt"
)

// solveState keeps computed state for the solver algorithm.
type solveState struct {

	// contains ascii digits (solved cells) and zero bytes (unsolved cells)
	board [81]byte

	// number of solved cells
	solvedCount int

	// eliminated[cell][digit] == true means that digit is not a solution
	// for cell
	eliminated [81][9]bool
}

// noSolutionErr is returned when no solutions were found
var noSolutionErr = fmt.Errorf("no solution")

// Solve the given sudoku. The given sudoku must consist of ascii digits (0-9)
// and dots(.) and the length must be 81. The solution will be returned as a
// string of ascii digits of length 81. An error will be returned if the sudoku
// is not solvable or if there is more than one valid soution.
func Solve(sudoku string) (string, error) {
	if len(sudoku) != 81 {
		return "", fmt.Errorf("illegal line length: %d", len(sudoku))
	}

	var st solveState
	for i, c := range sudoku {
		if c != '.' {
			st.solveCell(i, byte(c))
		}
	}

	return st.solve()
}

// solve the sudoku. See func Solve.
func (st *solveState) solve() (string, error) {
	for st.solvedCount != 81 {

		if st.solveDigitForCell() != 0 {
			continue
		}

		if st.solveCellForDigit() != 0 {
			continue
		}

		return st.search()
	}
	return string(st.board[:]), nil
}

// solveCell sets the digit for the given cell.
func (st *solveState) solveCell(cell int, digit byte) {
	if st.board[cell] != 0 {
		panic("cell already solved")
	}

	st.solvedCount++
	st.board[cell] = digit

	for _, cell := range neighbours[cell] {
		st.eliminated[cell][digit-'1'] = true
	}

	for digit := 0; digit < 9; digit++ {
		st.eliminated[cell][digit] = true
	}
}

// solveCellForDigit searches every group to find a digit that can only be in
// one cell. Solve the cells that are found and return the number of cells that
// were solved.
func (st *solveState) solveCellForDigit() int {
	var res int
	for _, cells := range cellGroups {
	outer:
		for digit := 0; digit < 9; digit++ {
			cell := -1
			for _, c := range cells {
				if st.eliminated[c][digit] {
					continue
				}

				if cell != -1 {
					continue outer
				}
				cell = c
			}
			if cell != -1 {
				st.solveCell(cell, byte('1'+digit))
				res++
			}
		}
	}
	return res
}

// solveDigitForCell searches the board for cells that have only one possible
// digit. Solve the cells that are found and return the number of cells that
// were solved.
func (st *solveState) solveDigitForCell() int {
	var res int
outer:
	for cell, x := range st.board {
		if x != 0 {
			continue //already solved
		}

		var digit byte
		for d, eliminated := range st.eliminated[cell] {
			if eliminated {
				continue
			}

			if digit != 0 {
				continue outer // multiple alternatives
			}
			digit = byte('1' + d)
		}

		if digit != 0 {
			st.solveCell(cell, digit)
			res++
		}
	}
	return res
}

// search for a solution. This function is expensive and is invoked only when
// other strategies fail to make progress. A solution is found by picking a
// cell to solve, then attempt every valid digit for that cell. Each attempt
// is a recursive call to solveState.solve() with a new solveState.
func (st *solveState) search() (string, error) {
	minAlts := 10
	// Select which selectedCell to search. Find unsolved selectedCell with the smallest
	// number of valid alternatives to minimize the search space.
	var selectedCell int
outer:
	for cell, digit := range st.board {
		if digit != 0 {
			continue //already solved
		}

		var alts int
		for _, eliminated := range st.eliminated[cell] {
			if eliminated {
				continue
			}

			alts++
			if alts >= minAlts {
				continue outer
			}
		}

		if alts == 0 {
			return "", noSolutionErr
		}

		selectedCell = cell

		if alts == 2 {
			break // because 2 is the optimal number of alternatives
		}

		minAlts = alts
	}

	var solution string

	for digit, eliminated := range st.eliminated[selectedCell] {
		if eliminated {
			continue
		}

		var attempt solveState = *st
		attempt.solveCell(selectedCell, byte('1'+digit))
		a, err := attempt.solve()

		if err == noSolutionErr {
			continue
		}

		if err != nil {
			return a, err
		}

		// a is a valid solution! Continue the loop only to detect if
		// there are more solutions.

		if solution != "" {
			return "", fmt.Errorf("multiple solutions")
		}
		solution = a
	}

	if solution == "" {
		return "", noSolutionErr
	}

	return solution, nil
}
