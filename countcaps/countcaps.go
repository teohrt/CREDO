package countcaps

import (
	"encoding/json"
	"net/http"
	"unicode"

	"github.com/gorilla/mux"
)

// CountCaps takes a string, and returns number of capital letters
func CountCaps(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsUpper(char) {
			count++
		}
	}
	return count
}

// CountCapsHandler returns the number capitals in the word passed in parameters
func CountCapsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	numOfCaps := CountCaps(params["word"])
	json.NewEncoder(w).Encode(numOfCaps)
}
