package documents

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type ReviewModel struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Title      string        `bson:"title,omitempty"`
	Body       string        `bson:"body,omitempty"`
	Rate       int           `bson:"rate,omitempty"`
	Files      []string      `bson:"files,omitempty"`
	Reviewer   *Reviewer     `bson:"reviewer,omitempty"`
	ReviewedAt time.Time     `bson:"reviewed_at,omitempty"`
}

type Reviewer struct {
	ID   string `bson:"id,omitempty"`
	Name string `bson:"name,omitempty"`
	Icon string `bson:"icon,omitempty"`
}
