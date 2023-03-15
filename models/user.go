package models

import (
	"time"
)

//if not declared, set default value / ex> rt_token = "", time = "0.0.000.000."
type User struct {
	Email     string
	Password  string
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	RtToken   string    `bson:"rt_token"`
}