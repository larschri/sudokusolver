// dimensions.go provides two static lookup tables and initalization code.

package solver

var (
	// cellGroups is a static lookup table that contains 27 cellGroups (9 rows + 9
	// columns + 9 "squares"). Each group contain the indices of 9 member
	// cells.
	cellGroups [27][9]int

	// neighbours is a static lookup table that contains the neighbours for
	// each cell (81 cells). Two cells are considered neighbours if they
	// are members of the same group (column, row or "square"). Each cell
	// has 20 neighbours.
	neighbours [81][20]int
)

// groupify is a helper function to initialize groups
func groupify(builder [][]int, s string) {
	for i, cc := range s {
		c := int(cc - '0')
		builder[c] = append(builder[c], i)
	}
}

// appendUnique is similar to the append builtin function, but does nothing if
// v is already in s. Used only for initialization.
func appendUnique(s []int, v int) []int {
	for _, n := range s {
		if n == v {
			return s
		}
	}
	return append(s, v)
}

func init() {
	// builder populates the pre-allocated arrays
	var builder [][]int
	for i := range cellGroups {
		builder = append(builder, cellGroups[i][:0])
	}
	groupify(builder[0:9], "000000000111111111222222222333333333444444444555555555666666666777777777888888888")
	groupify(builder[9:18], "012345678012345678012345678012345678012345678012345678012345678012345678012345678")
	groupify(builder[18:27], "000111222000111222000111222333444555333444555333444555666777888666777888666777888")

	builder = nil
	for i := range neighbours {
		builder = append(builder, neighbours[i][:0])
	}
	for _, cells := range cellGroups {
		for _, cell := range cells {
			for _, cell2 := range cells {
				if cell != cell2 {
					builder[cell] = appendUnique(builder[cell], cell2)
				}
			}
		}
	}
}
