package handler

import (
	"justvesit/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func StartItBuddy(w http.ResponseWriter, rf *http.Request) {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods("POST")

	http.ListenAndServe(":4000", r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// database.Init()
	// registeredusers.InitializeRegisteredUsers()
	controller.Home(w, r)
}
