package server

import (
	"fmt"
	"justvesit/router"
	"net/http"
)

func StartTheServerOnPort(port string) {

	r := router.Router()
	fmt.Println("Server is getting started...")
	fmt.Println("Server listening on ", port)
	http.ListenAndServe(port, r)
}
