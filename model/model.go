package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Email   string             `json:"email,omitempty"`
	Subject string             `json:"subject,omitempty"`
	Message string             `json:"message,omitempty"`
}

type WatchStamp struct {
	ID        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	WatchTime primitive.Timestamp `json:"watchTime,omitempty"`
}
