package model

import (
	"gopkg.in/mgo.v2/bson"
)

//UserModel is useing of user data
type UserModel struct {
	UserName      string `json:"user_name" bson:"user_name"`
	EmailAddress  string `json:"email_address" bson:"email_address"`
	ContactNumber string `json:"contact_number" bson:"contact_number"`
	Token         string `json:"token" bson:"token,omitempty" `
}

//UserModelID is useing of user data
type UserModelID struct {
	UserName      string        `json:"user_name" bson:"user_name"`
	EmailAddress  string        `json:"email_address" bson:"email_address"`
	ContactNumber string        `json:"contact_number" bson:"contact_number"`
	Token         string        `json:"token" bson:"token,omitempty" `
	ID            bson.ObjectId `json:"_id" bson:"_id"`
}
