package main

import (
	"justvesit/database"
	registeredusers "justvesit/registered_users"
	"justvesit/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DB_STRING")
	database.Init(connectionString)
	registeredusers.InitializeRegisteredUsers()
	server.StartTheServerOnPort(":4000")
}
