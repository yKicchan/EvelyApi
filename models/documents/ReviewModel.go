package documents

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type ReviewModel struct {
	ID         bson.ObjectId `bson:"_id"`
	Title      string        `bson:"title"`
	Body       string        `bson:"body"`
	Rate       int           `bson:"rate"`
	Files []string `bson:"files"`
	Reviewer   *Reviewer     `bson:"reviewer"`
	ReviewedAt time.Time     `bson:"reviewed_at"`
}

type Reviewer struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
	Icon string `bson:"icon"`
}
