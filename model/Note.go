package model

import "gopkg.in/mgo.v2/bson"

// Note - Struct notes about a residency
type Note struct {
	ID          bson.ObjectId `bson:"Id"          json:"Id"`
	ResidencyID string        `bson:"residencyId" json:"residencyId"`
	Notes       string        `bson:"notes"       json:"notes"`
}
