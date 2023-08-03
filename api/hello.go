package hello

import (
	"fmt"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		currentTime := time.Now().Format(time.RFC850)
		fmt.Fprint(w, currentTime)
	} else if r.Method == http.MethodPost {
		fmt.Fprint(w, "Received a POST request to /api/date")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
