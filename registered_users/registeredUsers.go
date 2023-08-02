package registeredusers

import (
	"fmt"
	"justvesit/models"
)

func InitializeRegisteredUsers() {
	fmt.Println("initializing registered users")
	userList := []models.User{}

	user1 := models.User{UserName: "user1", Password: "password1"}

	userList = append(userList, user1)
	fmt.Println(userList)
}
