package repository

import (
	"log"

	"gopkg.in/mgo.v2"
)

type RConn struct{}

var db *mgo.Database

func (r *RConn) connect(Server string, Database string) {
	session, err := mgo.Dial(Server)

	// Verify if occurs some error and stop application
	if err != nil {
		log.Fatal(err)
	}

	// Establish a session with database
	db = session.DB(Database)
}
