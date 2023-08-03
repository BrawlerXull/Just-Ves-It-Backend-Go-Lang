package main

import (
	"justvesit/database"
	registeredusers "justvesit/registered_users"
	"justvesit/server"
)

func Main() {
	database.Init()
	registeredusers.InitializeRegisteredUsers()
	server.StartTheServerOnPort(":4000")
}
