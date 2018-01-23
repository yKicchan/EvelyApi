package models

import (
	. "EvelyApi/config"
	. "EvelyApi/models/collections"
	"labix.org/v2/mgo"
)

/**
 * Evelyのデーターベースを操作するためのオブジェクト
 */
type EvelyDB struct {
	*mgo.Database
	Events  *EventsCollection
	Users   *UsersCollection
	Reviews *ReviewsCollection
}

func NewEvelyDB(db *mgo.Database) *EvelyDB {
	return &EvelyDB{
		db,
		NewEventsCollection(db.C(EVENTS_COLLECTION)),
		NewUsersCollection(db.C(USERS_COLLECTION)),
		NewReviewsCollection(db.C(REVIEWS_COLLECTION)),
	}
}
