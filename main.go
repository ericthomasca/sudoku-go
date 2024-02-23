package main

import (
	"encoding/json"
	"net/http"

	"github.com/ericthomasca/sudoku-go-api/sudoku"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/solve", solveHandler)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8976", handler)

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

	if puzzle.Solve() {
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

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	} else {
		http.Error(w, "No solution exists", http.StatusNotFound)
	}
}
