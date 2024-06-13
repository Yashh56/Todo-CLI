package model

import "time"

type Model struct {
	Title     string    `json:"title" bson:"title"`
	Done      bool      `json:"done" bson:"done"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
