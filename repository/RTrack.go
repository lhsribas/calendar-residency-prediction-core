package repository

import (
	"github.com/lhsribas/residency-prediction-backend/model"
	"gopkg.in/mgo.v2/bson"
)

type RTrack struct{}

const (
	TrackCOLL = "track"
)

func (r *RTrack) saveT(m model.Track) error {
	return db.C(TrackCOLL).Insert(&m)
}

func (r *RTrack) findByKeyT(Id string) ([]model.Track, error) {
	var m []model.Track

	err := db.C(TrackCOLL).Find(bson.M{"key": Id}).All(&m)
	return m, err
}
