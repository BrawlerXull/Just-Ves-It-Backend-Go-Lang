package controller

import (
	"encoding/json"
	"fmt"
	registeredusers "justvesit/registered_users"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to home")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello"})
	return
}

func ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := registeredusers.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}
