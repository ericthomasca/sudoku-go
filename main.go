package main

import (
	"fmt"

	"github.com/ericthomasca/sudoku-go/sudoku"
)

type SudokuPuzzle [9][9]int

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

func NewSudokuPuzzle(grid [9][9]int) SudokuPuzzle {
    return SudokuPuzzle(grid)
}

func (p SudokuPuzzle) Print() {
    for _, row := range p {
        fmt.Println(row)
    }
}

func (p SudokuPuzzle) FindEmptyLocation() (int, int, bool) {
    for row := range p {
        for col := range p[row] {
            if p[row][col] == 0 {
                return row, col, true
            }
        }
    }
    return -1, -1, false
}

func (p SudokuPuzzle) IsSafe(row, col, num int) bool {
    return !p.usedInRow(row, num) && !p.usedInCol(col, num) && !p.usedInBox(row-row%3, col-col%3, num)
}

func (p SudokuPuzzle) usedInRow(row, num int) bool {
    for _, value := range p[row] {
        if value == num {
            return true
        }
    }
    return false
}

func (p SudokuPuzzle) usedInCol(col, num int) bool {
    for _, row := range p {
        if row[col] == num {
            return true
        }
    }
    return false
}

func (p SudokuPuzzle) usedInBox(boxStartRow, boxStartCol, num int) bool {
    for row := 0; row < 3; row++ {
        for col := 0; col < 3; col++ {
            if p[row+boxStartRow][col+boxStartCol] == num {
                return true
            }
        }
    }
    return false
}

func (p *SudokuPuzzle) Solve() bool {
    row, col, found := p.FindEmptyLocation()
    if !found {
        return true
    }
    for num := 1; num <= 9; num++ {
        if p.IsSafe(row, col, num) {
            p[row][col] = num
            if p.Solve() {
                return true
            }
            p[row][col] = 0
        }
    }
    return false
}


