package sudoku

import "fmt"

type Puzzle [9][9]int

func New(grid [9][9]int) Puzzle {
	return Puzzle(grid)
}

func (p Puzzle) Print() {
	for _, row := range p {
		fmt.Println(row)
	}
}

func (p *Puzzle) Solve() bool {
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

func (p Puzzle) IsSafe(row, col, num int) bool {
	return !p.usedInRow(row, num) && !p.usedInCol(col, num) && !p.usedInBox(row-row%3, col-col%3, num)
}

func (p Puzzle) usedInRow(row, num int) bool {
	for _, value := range p[row] {
		if value == num {
			return true
		}
	}
	return false
}

func (p Puzzle) usedInCol(col, num int) bool {
	for _, row := range p {
		if row[col] == num {
			return true
		}
	}
	return false
}

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
