package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Team - Struct to refers the team that can work in a residency
type StaffMember struct {
	ID          bson.ObjectId `bson:"_Id"           json:"Id"`
	Name        string        `bson:"name"          json:"name"`
	Email       string        `bson:"email"         json:"email"`
	Photo       int64         `bson:"photo"         json:"photo"`
	Core        bool          `bson:"isCore"        json:"isCore"`
	ResidencyID string        `bson:"residencyId"   json:"residencyId"`
}
