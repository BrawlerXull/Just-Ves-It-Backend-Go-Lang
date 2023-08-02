package registeredusers

import (
	"fmt"
	"justvesit/models"
)

var userList []models.User

func InitializeRegisteredUsers() {
	fmt.Println("initializing registered users")
	user1 := models.User{UserName: "user1", Password: "password1"}
	userList = append(userList, user1)

	fmt.Println(userList)
}

func GetAllUsers() []models.User {
	return userList
}
