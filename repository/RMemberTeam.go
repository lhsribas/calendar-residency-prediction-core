package repository

import (
	"github.com/lhsribas/calendar-residency-prediction-core/model"
	"gopkg.in/mgo.v2/bson"
)

type RMemberTeam struct{}

const (
	MemberTeamCOLL = "member_team"
)

func (r *RMemberTeam) saveMT(m model.MemberTeam) error {
	return db.C(TrackCOLL).Insert(&m)
}

func (r *RMemberTeam) findByKeyMT(Id string) ([]model.MemberTeam, error) {
	var m []model.MemberTeam

	err := db.C(MemberTeamCOLL).Find(bson.M{"residencyId": Id}).All(&m)
	return m, err
}
