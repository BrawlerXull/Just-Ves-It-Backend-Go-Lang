package main

import (
	"fmt"
	"justvesit/database"
	registeredusers "justvesit/registered_users"
	"justvesit/router"
	"net/http"
)

func main() {
	database.Init()
	registeredusers.InitializeRegisteredUsers()
	fmt.Println("Hello welcome to the world!")
	r := router.Router()
	fmt.Println("Server is getting started...")
	http.ListenAndServe(":4000", r)
	println("Server listening on :4000")
}
