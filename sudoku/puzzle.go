package sudoku

import "fmt"

// Puzzle represents a Sudoku puzzle grid.
type Puzzle [9][9]int

// New creates a new Sudoku puzzle from the given grid.
func New(grid [9][9]int) Puzzle {
	return Puzzle(grid)
}

// Print prints the Sudoku puzzle grid.
func (p Puzzle) Print() {
	for _, row := range p {
		fmt.Println(row)
	}
}

// Solve solves the Sudoku puzzle using backtracking algorithm.
func (p Puzzle) Solve() bool {
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

// FindEmptyLocation finds an empty cell in the Sudoku puzzle.
// It returns the row and column indices of the empty cell and a boolean indicating if an empty cell was found.
func (p Puzzle) FindEmptyLocation() (int, int, bool) {
	for row := range p {
		for col := range p[row] {
			if p[row][col] == 0 {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}

// IsSafe checks if it's safe to place a number in the given row and column of the Sudoku puzzle.
func (p Puzzle) IsSafe(row, col, num int) bool {
	return !p.usedInRow(row, num) && !p.usedInCol(col, num) && !p.usedInBox(row-row%3, col-col%3, num)
}

// usedInRow checks if the given number is already used in the specified row of the Sudoku puzzle.
func (p Puzzle) usedInRow(row, num int) bool {
	for _, value := range p[row] {
		if value == num {
			return true
		}
	}
	return false
}

// usedInCol checks if the given number is already used in the specified column of the Sudoku puzzle.
func (p Puzzle) usedInCol(col, num int) bool {
	for _, row := range p {
		if row[col] == num {
			return true
		}
	}
	return false
}

// usedInBox checks if the given number is already used in the 3x3 box containing the cell at the specified row and column.
func (p Puzzle) usedInBox(boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if p[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}
