package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("welcome to home")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello"})
		return
	} else if r.Method == http.MethodPost {
		fmt.Fprint(w, "Received a POST request to /api/date")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
