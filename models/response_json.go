package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ErrResponse struct {
	Message interface{} `json:"message"`
}

type OkResponse struct {
	User    primitive.M   `json:"user,omitempty"`
	Users   []primitive.M `json:"users,omitempty"`
	Message string        `json:"message"`
}

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message"`
}
