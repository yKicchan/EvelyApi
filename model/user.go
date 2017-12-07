package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDB struct {
	*mgo.Database
}

type UserModel struct {
	ID       string `bson:id`
	Password string `bson:password`
	Name     string `bson:name`
	Mail     string `bson:mail`
	Tel      string `bson:tel`
}

func NewUserDB(db *mgo.Database) *UserDB {
	return &UserDB{db}
}

/**
 * ログイン認証する
 */
func (db *UserDB) Authentication(id string, password string) (user *UserModel, err error) {
	query := bson.M{"id": id, "password": password}
	err = db.C("users").Find(query).One(&user)
	return
}
