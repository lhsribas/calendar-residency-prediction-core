package repository

import (
	"github.com/lhsribas/residency-prediction-backend/model"
	"gopkg.in/mgo.v2/bson"
)

type RStaffMember struct{}

const (
	MemberTeamCOLL = "staff_member"
)

func (r *RStaffMember) saveSM(m model.StaffMember) error {
	return db.C(TrackCOLL).Insert(&m)
}

func (r *RStaffMember) findByKeySM(Id string) ([]model.StaffMember, error) {
	var m []model.StaffMember

	err := db.C(MemberTeamCOLL).Find(bson.M{"residencyId": Id}).All(&m)
	return m, err
}
