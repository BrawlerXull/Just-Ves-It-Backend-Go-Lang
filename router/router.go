package router

import (
	"justvesit/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Home).Methods("GET")
	router.HandleFunc("/users", controller.ShowAllUsers).Methods("GET")
	router.HandleFunc("/auth", controller.AuthenticateTheUser).Methods("POST")
	router.HandleFunc("/all", controller.GetAllMyTasks).Methods("GET")
	return router

}
