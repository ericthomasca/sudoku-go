package main

import (
	"fmt"

	"github.com/ericthomasca/sudoku-go/sudoku"
)

func main() {
	grid := [9][9]int{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}

	puzzle := sudoku.New(grid)

	fmt.Println("Unsolved Sudoku Puzzle:")
	puzzle.Print()

	if puzzle.Solve() {
		fmt.Println("\nSolved Sudoku Puzzle:")
		puzzle.Print()
	} else {
		fmt.Println("No solution exists")
	}
}
