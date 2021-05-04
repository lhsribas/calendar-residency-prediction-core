package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Track struct {
	ID        bson.ObjectId `bson:"_Id"           json:"Id"`
	Key       string        `bson:"key"           json:"key"`
	CreatedAt time.Time     `bson:"CreatedAt"     json:"CreatedAt"`
	UpdatedAt time.Time     `bson:"UpdatedAt"     json:"UpdatedAt"`
	UpdatedBy string        `bson:"UpdatedBy"     json:"UpdatedBy"`
}
