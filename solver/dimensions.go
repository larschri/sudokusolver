package solver

var (
	groups     [27][9]int
	neighbours [81][20]int
)

func groupify(builder [][]int, s string) {
	for i, cc := range s {
		c := int(cc - '0')
		builder[c] = append(builder[c], i)
	}
}

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
	for i := range groups {
		builder = append(builder, groups[i][:0])
	}
	groupify(builder[0:9], "000000000111111111222222222333333333444444444555555555666666666777777777888888888")
	groupify(builder[9:18], "012345678012345678012345678012345678012345678012345678012345678012345678012345678")
	groupify(builder[18:27], "000111222000111222000111222333444555333444555333444555666777888666777888666777888")

	builder = nil
	for i := range neighbours {
		builder = append(builder, neighbours[i][:0])
	}
	for _, grp := range groups {
		for _, pos := range grp {
			for _, pos2 := range grp {
				if pos != pos2 {
					builder[pos] = appendUnique(builder[pos], pos2)
				}
			}
		}
	}
}
