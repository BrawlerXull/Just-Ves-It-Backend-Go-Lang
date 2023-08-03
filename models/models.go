package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	ID          primitive.ObjectID `json:"_id"`
	Date        time.Time          `json:"date"`
	Description string             `json:"description"`
	Subject     string             `json:"subject"`
}
