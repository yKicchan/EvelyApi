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
	e := &EventModel{
		ID:        id,
		Image:     p.Image,
		Title:     p.Title,
		Body:      p.Body,
		Files:     p.Files,
		Categorys: p.Categorys,
		Host: &Host{
			ID:   u.ID,
			Name: u.Name,
			Icon: u.Icon,
		},
		Schedules:   toSchedulesModel(p.Schedules),
		NoticeRange: p.NoticeRange,
		Scope:       p.Scope,
		OpenFlg:     p.OpenFlg,
		UpdateDate:  time.Now(),
	}
	if p.Mail != nil {
		e.Mail = *p.Mail
	}
	if p.Tel != nil {
		e.Tel = *p.Tel
	}
	if p.URL != nil {
		e.URL = *p.URL
	}
	return e
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
			Icon: u.Icon,
		},
	}
}
