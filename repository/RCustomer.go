package repository

import (
	"github.com/lhsribas/calendar-residency-prediction-core/model"
	"gopkg.in/mgo.v2/bson"
)

type RCustomer struct{}

const (
	CustomerCOLL = "customer"
)

func (r *RCustomer) saveC(m model.Customer) error {
	return db.C(CustomerCOLL).Insert(&m)
}

func (r *RCustomer) fundByIdC(Id string) (model.Customer, error) {
	var m model.Customer

	err := db.C(CustomerCOLL).FindId(bson.ObjectIdHex(Id)).One(&m)
	return m, err
}
