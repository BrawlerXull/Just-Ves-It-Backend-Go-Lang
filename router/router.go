package router

import (
	"github.com/gorilla/mux"

	"justvesit/methods"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", methods.Home).Methods("GET")
	return router

}
