package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"justvesit/database"
	"justvesit/models"
	registeredusers "justvesit/registered_users"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to home")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello"})
	return
}

func ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := registeredusers.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func AuthenticateTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var incomingUserData models.User
	err := json.NewDecoder(r.Body).Decode(&incomingUserData)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return
	}

	for _, u := range registeredusers.GetAllUsers() {
		if incomingUserData.UserName == u.UserName && incomingUserData.Password == u.Password {
			json.NewEncoder(w).Encode(map[string]interface{}{"resp": u})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{})
}

func GetAllTasks() []primitive.M {
	cur, err := database.Collection().Find(context.Background(), bson.D{{}})
	checkError(err)
	var tasks []primitive.M

	for cur.Next(context.Background()) {
		var task bson.M
		err := cur.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	defer cur.Close(context.Background())
	return tasks
}
func GetAllMyTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allTasks := GetAllTasks()
	json.NewEncoder(w).Encode(allTasks)
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func SaveMyTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var task models.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	SaveTask(task)
	json.NewEncoder(w).Encode(task)

}
func SaveTask(task models.Task) {

	inserted, err := database.Collection().InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)

}
