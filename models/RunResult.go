package models

import (
	. "EvelyApi/models/documents"
)

type RunResult struct {
	Results []*Result `bson:"results"`
	Status  *Status   `bson:"status"`
	Ok      int       `bson:"ok"`
}

type Result struct {
	Distance float64     `bson:"dis"`
	Event    *EventModel `bson:"obj"`
}

type Status struct {
	Nscanned      interface{} `bson:"nscanned"`
	ObjectsLoaded interface{} `bson:"objectsLoaded"`
	AvgDistance   float64     `bson:"avgDistance"`
	MaxDistance   float64     `bson:"maxDistance"`
	Time          int         `bson:"time"`
}
