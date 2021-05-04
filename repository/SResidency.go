package repository

import (
	"time"

	"github.com/lhsribas/calendar-residency-prediction-core/model"
	"gopkg.in/mgo.v2/bson"
)

type SResidency struct {
	Server   string
	Database string
}

var rConn = RConn{}

var rCustomer = RCustomer{}
var rStaffMember = RStaffMember{}
var rNote = RNote{}
var rResindecy = RResindecy{}
var rTrack = RTrack{}

func (s *SResidency) Connect() {
	rConn.connect(s.Server, s.Database)
}

// Customer

func (s *SResidency) SaveCustomer(m model.Customer, userSystem string) (bson.ObjectId, error) {
	m.ID = bson.NewObjectId()

	// creates a track this method
	createTrack(m.ID.Hex(), userSystem)

	return m.ID, rCustomer.saveC(m)
}

func (s *SResidency) FindCustomerById(Id string) (model.Customer, error) {
	return rCustomer.fundByIdC(Id)
}

// Residency

func (s *SResidency) CreateResidencyIndex() {
	rResindecy.residencyIndex()
}

func (s *SResidency) SaveResidency(m model.Residency, userSystem string) (bson.ObjectId, error) {
	m.ID = bson.NewObjectId()

	// creates a track this method
	createTrack(m.ID.Hex(), userSystem)

	return m.ID, rResindecy.saveR(m)
}

func (s *SResidency) FindResidencyById(Id string) (model.Residency, error) {
	return rResindecy.findByIdR(Id)
}

func (s *SResidency) FindResidencyBySalesKey(salesKey string) (model.Residency, error) {
	return rResindecy.findByKeyR(salesKey)
}

// Member Team

func (s *SResidency) SaveStaffMember(m model.StaffMember, userSystem string) (bson.ObjectId, error) {
	m.ID = bson.NewObjectId()

	// creates a track this method
	createTrack(m.ID.Hex(), userSystem)

	return m.ID, rStaffMember.saveSM(m)
}

func (S *SResidency) FindStaffMembersByResidencyId(Id string) ([]model.StaffMember, error) {
	return rStaffMember.findByKeySM(Id)
}

// Note

func (s *SResidency) SaveResidencyNotes(m model.Note, userSystem string) (bson.ObjectId, error) {
	m.ID = bson.NewObjectId()

	// creates a track this method
	createTrack(m.ID.Hex(), userSystem)

	return m.ID, rNote.saveN(m)
}

func (s *SResidency) FindResidencyNotes(residencyId string) ([]model.Note, error) {
	return rNote.findByResidencyId(residencyId)
}

func createTrack(Id string, userSystem string) error {
	track := model.Track{
		ID:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedBy: userSystem,
		Key:       Id,
		UpdatedAt: time.Now(),
	}

	return rTrack.saveT(track)
}
