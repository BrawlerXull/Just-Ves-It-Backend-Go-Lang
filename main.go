package main

import (
	"fmt"
	"justvesit/database"
	"justvesit/router"
	"net/http"
)

func main() {
	database.Init()
	fmt.Println("Hello welcome to the world!")
	r := router.Router()
	fmt.Println("Server is getting started...")
	http.ListenAndServe(":4000", r)
	println("Server listening on :4000")
}
