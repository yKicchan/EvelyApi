package findOptions

import (
    "labix.org/v2/mgo/bson"
)

// イベントIDで検索するときのオプション
type EventIDOption interface {
    FindOptions
    SetEventID(bson.ObjectId)
    GetEventID() bson.ObjectId
    IsEventIDSet() bool
}
