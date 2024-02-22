package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ericthomasca/sudoku-go-api/sudoku"
)

func main() {
	http.HandleFunc("/solve", solveHandler)
	fmt.Println("Listening on port 8976...")
	http.ListenAndServe(":8976", nil)

}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input struct {
        Puzzle [9][9]int `json:"puzzle"`
    }

    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	puzzle := sudoku.New(input.Puzzle)

	fmt.Println("Unsolved Sudoku Puzzle:")
	puzzle.Print()

	if puzzle.Solve() {
		fmt.Println("\nSolved Sudoku Puzzle:")
		puzzle.Print()

		response := struct {
			SolvedPuzzle sudoku.Puzzle `json:"solved_puzzle"`
		}{
			SolvedPuzzle: puzzle,
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	} else {
		fmt.Println("No solution exists")

        http.Error(w, "No solution exists", http.StatusNotFound)
	}
}
