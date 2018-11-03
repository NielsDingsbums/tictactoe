package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
)

// Global vars
var state GameState = GameState{[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}}

func GetState(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("\n=== GET  REQUEST ===")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://snap.berkeley.edu")

	params := mux.Vars(r)

	cell, _ := strconv.Atoi(params["id"])
	fmt.Printf("GET Cell %d: %d\n", cell, state.cells[cell])

	_ = json.NewEncoder(w).Encode(state.cells[cell])
}

func UpdateState(w http.ResponseWriter, r *http.Request) {

	var data [9]int
	_ = json.NewDecoder(r.Body).Decode(&data)
	state.cells = data
	fmt.Println("\n===  POST REQUEST  ===")

	for i, cell := range state.cells {
		if math.Mod(float64(i)+1.0, 3.0) == 0 {
			switch cell {
			case 0:
				fmt.Println("- ")
			case 1:
				fmt.Println("X ")
			case 2:
				fmt.Println("O ")
			}
		} else {
			switch cell {
			case 0:
				fmt.Print("- ")
			case 1:
				fmt.Print("X ")
			case 2:
				fmt.Print("O ")
			}
		}
	}
	fmt.Println("")

}
