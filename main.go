package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		if len(bytes) != 81 {
			panic("Illegal line length: " + string(bytes))
		}

		var s solver
		for i := 0; i < 81; i++ {
			if bytes[i] != '.' {
				s.setSolved(i, bytes[i])
			}
		}

		if err := s.solve(); err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(string(s.board[:]))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func prettyBoard(b [81]byte) string {
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
			board[i] = b[j]
			j++
		}
	}

	return string(board)
}
