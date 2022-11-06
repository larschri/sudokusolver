package main

type (
	group [9]int

	dimension struct {
		groups [9]group
		idxs   [81]int
	}
)

var (
	rows   dimension
	cols   dimension
	boxes  dimension
	groups [27]group
)

func (d *dimension) init(board [9]string) {
	var slices [9][]int
	var idxs []int
	for i := 0; i < 81; i++ {
		c := int(board[i/9][i%9] - '0')
		slices[c] = append(slices[c], i)
		idxs = append(idxs, c)
	}
	for i := 0; i < 9; i++ {
		if copy(d.groups[i][:], slices[i]) != 9 {
			panic("urk")
		}
	}

	if copy(d.idxs[:], idxs) != 81 {
		panic("burk")
	}
}

func init() {
	rows.init([9]string{
		"000000000",
		"111111111",
		"222222222",
		"333333333",
		"444444444",
		"555555555",
		"666666666",
		"777777777",
		"888888888",
	})
	cols.init([9]string{
		"012345678",
		"012345678",
		"012345678",
		"012345678",
		"012345678",
		"012345678",
		"012345678",
		"012345678",
		"012345678",
	})
	boxes.init([9]string{
		"000111222",
		"000111222",
		"000111222",
		"333444555",
		"333444555",
		"333444555",
		"666777888",
		"666777888",
		"666777888",
	})
	copy(groups[0:9], rows.groups[:])
	copy(groups[9:18], cols.groups[:])
	copy(groups[18:27], boxes.groups[:])
}
