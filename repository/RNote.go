package repository

import (
	"github.com/lhsribas/calendar-residency-prediction-core/model"
	"gopkg.in/mgo.v2/bson"
)

type RNote struct{}

const (
	NoteCOLL = "notes"
)

func (r *RNote) saveN(m model.Note) error {
	return db.C(NoteCOLL).Insert(m)
}

func (r *RNote) findByResidencyId(Id string) ([]model.Note, error) {
	var m []model.Note

	err := db.C(NoteCOLL).Find(bson.M{"residencyId": Id}).All(&m)
	return m, err
}
