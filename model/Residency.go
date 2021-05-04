package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Residency - Struct a resdency that have a date predction to start
type Residency struct {
	ID          bson.ObjectId `bson:"_Id"          json:"Id"`
	Name        string        `bson:"name"         json:"name"`
	SalesKey    string        `bson:"salesKey"     json:"salesKey"`
	Description string        `bson:"description"  json:"description"`
	Prediction  time.Time     `bson:"prediction"   json:"prediction"`
	CustomerID  string        `bson:"customerId"   json:"customerId"`
}
