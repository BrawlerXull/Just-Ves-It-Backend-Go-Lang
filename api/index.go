package main

import (
	"justvesit/database"
	registeredusers "justvesit/registered_users"
	"justvesit/server"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	database.Init()
	registeredusers.InitializeRegisteredUsers()
	server.StartTheServerOnPort(":4000")
}

// func main() {

// }
