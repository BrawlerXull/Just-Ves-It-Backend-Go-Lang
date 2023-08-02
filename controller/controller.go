package controller

import (
	"encoding/json"
	"fmt"
	"justvesit/models"
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

func AuthenticateTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var incomingUserData models.User
	err := json.NewDecoder(r.Body).Decode(&incomingUserData)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	for _, u := range registeredusers.GetAllUsers() {
		if incomingUserData.UserName == u.UserName && incomingUserData.Password == u.Password {
			json.NewEncoder(w).Encode(map[string]interface{}{"resp": u})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{})
}
