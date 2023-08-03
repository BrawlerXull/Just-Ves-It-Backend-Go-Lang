package api

import (
	"context"
	"encoding/json"
	"fmt"
	"justvesit/database"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

func All(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		users := GetAllTasks()
		json.NewEncoder(w).Encode(users)
		return
	} else if r.Method == http.MethodPost {
		fmt.Fprint(w, "Received a POST request to /api/date")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetAllTasks() []primitive.M {
	database.Init()
	cur, err := database.Collection().Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal("Error finding tasks:", err)
		return nil
	}
	defer cur.Close(context.Background())

	var tasks []primitive.M
	for cur.Next(context.Background()) {
		var task bson.M
		err := cur.Decode(&task)
		if err != nil {
			log.Fatal("Error decoding task:", err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}
