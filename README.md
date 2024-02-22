# Sudoku Solver

This is a simple Sudoku solver implemented in Go.

## Introduction

Sudoku Solver is a command-line application that allows you to input an unsolved Sudoku puzzle and solve it using a backtracking algorithm.

## Features

- Create Sudoku puzzles from a 9x9 grid.
- Solve Sudoku puzzles using a backtracking algorithm.
- Print both unsolved and solved Sudoku puzzles.

## Usage

To use Sudoku Solver, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/sudoku-solver.git
   ```

2. Navigate to the project directory:

   ```bash
   cd sudoku-solver
   ```

3. Run the `main.go` file:

   ```bash
   go run main.go
   ```

4. Follow the instructions to input an unsolved Sudoku puzzle.

5. The program will solve the puzzle and display the solution.

## Example

Here's an example of using Sudoku Solver:

```
$ go run main.go
Unsolved Sudoku Puzzle:
[8 0 0 0 0 0 0 0 0]
[0 0 3 6 0 0 0 0 0]
[0 7 0 0 9 0 2 0 0]
[0 5 0 0 0 7 0 0 0]
[0 0 0 0 4 5 7 0 0]
[0 0 0 1 0 0 0 3 0]
[0 0 1 0 0 0 0 6 8]
[0 0 8 5 0 0 0 1 0]
[0 9 0 0 0 0 4 0 0]

Solved Sudoku Puzzle:
[8 1 2 7 5 3 6 4 9]
[9 4 3 6 8 2 1 7 5]
[6 7 5 4 9 1 2 8 3]
[1 5 4 2 3 7 8 9 6]
[3 6 9 8 4 5 7 2 1]
[2 8 7 1 6 9 5 3 4]
[5 2 1 9 7 4 3 6 8]
[4 3 8 5 2 6 9 1 7]
[7 9 6 3 1 8 4 5 2]
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
