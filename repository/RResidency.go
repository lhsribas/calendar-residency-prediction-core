package repository

import (
	"github.com/lhsribas/residency-prediction-backend/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RResindecy struct{}

const (
	residencyCOLL = "residency"
)

func (repository *RResindecy) residencyIndex() {
	index := mgo.Index{
		Key:        []string{"salesKey"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := db.C(residencyCOLL).EnsureIndex(index); err != nil {
		//stop the execution
		panic(err)
	}
}

func (r *RResindecy) saveR(m model.Residency) error {
	return db.C(residencyCOLL).Insert(&m)
}

func (r *RResindecy) findByIdR(Id string) (model.Residency, error) {
	var m model.Residency

	err := db.C(residencyCOLL).FindId(bson.ObjectIdHex(Id)).One(&m)
	return m, err
}

func (r *RResindecy) findByKeyR(salesKey string) (model.Residency, error) {
	var m model.Residency

	err := db.C(residencyCOLL).Find(bson.M{"salesKey": salesKey}).One(&m)
	return m, err
}
