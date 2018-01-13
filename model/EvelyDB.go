package model

import (
	. "EvelyApi/config"
	. "EvelyApi/model/collection"
	"labix.org/v2/mgo"
)

/**
 * Evelyのデーターベースを操作するためのオブジェクト
 */
type EvelyDB struct {
	*mgo.Database
    Events *EventsCollection
    Users *UsersCollection
}

func NewEvelyDB(db *mgo.Database) *EvelyDB {
	return &EvelyDB{
        db,
        NewEventsCollection(db.C(EVENT_COLLECTION)),
        NewUsersCollection(db.C(USER_COLLECTION)),
    }
}
