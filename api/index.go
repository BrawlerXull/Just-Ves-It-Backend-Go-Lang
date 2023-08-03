package handler

import (
	"justvesit/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// database.Init()
	// registeredusers.InitializeRegisteredUsers()
	controller.Home(w, r)
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods("GET")

	http.ListenAndServe(":4000", r)
}
