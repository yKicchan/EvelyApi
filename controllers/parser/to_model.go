package parser

import (
	"EvelyApi/app"
	. "EvelyApi/models/documents"
	"labix.org/v2/mgo/bson"
	"time"
)

/**
 * イベント情報をDBモデルに変換する
 * @param  p           イベント情報の入ったPayload
 * @param  id          イベントID
 * @param  user        イベントのホスト情報
 * @return *EventModel DBモデルに変換したイベント情報
 */
func ToEventModel(p *app.EventPayload, id bson.ObjectId, u *UserModel) *EventModel {
	return &EventModel{
		ID:    id,
		Image: p.Image,
		Title: p.Title,
		Body:  p.Body,
		Files: p.Files,
		Host: &Host{
			ID:   u.ID,
			Name: u.Name,
		},
		Mail:        p.Mail,
		Tel:         p.Tel,
		URL:         p.URL,
		Schedules:   toSchedulesModel(p.Schedules),
		NoticeRange: p.NoticeRange,
		Scope:       p.Scope,
		OpenFlg:     p.OpenFlg,
		UpdateDate:  time.Now(),
	}
}

func toSchedulesModel(oldSchedules []*app.Schedule) (newSchedules []*Schedule) {
	for _, old := range oldSchedules {
		schedule := &Schedule{
			Location: &Location{
				Name:   old.Location.Name,
				LngLat: [2]float64{old.Location.Lng, old.Location.Lat},
			},
			UpcomingDate: &UpcomingDate{
				StartDate: old.UpcomingDate.StartDate,
				EndDate:   old.UpcomingDate.EndDate,
			},
		}
		newSchedules = append(newSchedules, schedule)
	}
	return newSchedules
}

func ToReviewModel(p *app.ReviewPayload, id bson.ObjectId, u *UserModel) *ReviewModel {
	return &ReviewModel{
		ID:    id,
		Title: p.Title,
		Body:  p.Body,
		Files: p.Files,
		Rate:  p.Rate,
		Reviewer: &Reviewer{
			ID:   u.ID,
			Name: u.Name,
		},
	}
}
