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
}

func NewEvelyDB(db *mgo.Database) *EvelyDB {
	return &EvelyDB{db}
}

func (this *EvelyDB) Events() *EventsCollection {
	return NewEventsCollection(this.C(EVENT_COLLECTION))
}

func (this *EvelyDB) Users() *UsersCollection {
	return NewUsersCollection(this.C(USER_COLLECTION))
}

func (this *EvelyDB) PendingUsers() *PendingUsersCollection {
	return NewPendingUsersCollection(this.C(PENDING_USER_COLLECTION))
}
