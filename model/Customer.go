package model

import "gopkg.in/mgo.v2/bson"

// Customer - Struct a customer for a residency
type Customer struct {
	ID          bson.ObjectId `bson:"Id"           json:"Id"`
	Name        string        `bson:"name"         json:"name"`
	Description string        `bson:"description"  json:"description"`
	Photo       int64         `bson:"photo"        json:"photo"`
}
