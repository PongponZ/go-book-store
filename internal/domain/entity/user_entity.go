package entity

import "time"

type User struct {
	ID        string    `json:"id"         bson:"_id"`
	Username  string    `json:"username"   bson:"username"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name"  bson:"last_name"`
	Email     string    `json:"email"      bson:"email"`
	Phone     string    `json:"phone"      bson:"phone"`
	Password  string    `json:"-"          bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
