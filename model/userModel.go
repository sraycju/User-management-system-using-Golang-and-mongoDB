package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty"`
	Dob string               `json:"dob,omitempty"`
	Address string           `json:"add,omitempty"`
	Description string       `json:"desc,omitempty"`
	Creation string          `json:"createdAt,omitempty"`
}