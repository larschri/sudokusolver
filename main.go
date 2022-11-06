package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/larschri/sudokusolver/solver"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		board := scanner.Text()

		solution, err := solver.Solve(board)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(solution)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
