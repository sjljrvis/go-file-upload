package models

import "gopkg.in/mgo.v2/bson"

// User - Represents a User, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Email string        `bson:"email" json:"email"`
	Token string        `bson:"token" json:"token"`
}
