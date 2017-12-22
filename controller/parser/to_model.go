package parser

import (
    "EvelyApi/app"
    "EvelyApi/model"
    "time"
)

/**
 * イベント情報をDBモデルに変換する
 * @param  p           イベント情報の入ったPayload
 * @param  id          イベントID
 * @param  user        イベントのホスト情報
 * @return *EventModel DBモデルに変換したイベント情報
 */
func ToEventModel(p *app.EventPayload, id string, user *model.UserModel) *model.EventModel {
	return &model.EventModel{
		ID:    id,
		Title: p.Title,
		Host: model.Host{
			ID:   user.ID,
			Name: user.Name,
		},
		Body: p.Body,
		Place: model.Location{
			Name:   p.Place.Name,
			LngLat: [2]float64{p.Place.Lng, p.Place.Lat},
		},
		UpdateDate: time.Now(),
		UpcomingDate: model.UpcomingDate{
			StartDate: p.UpcomingDate.StartDate,
			EndDate:   p.UpcomingDate.EndDate,
		},
		URL:  p.URL,
		Mail: p.Mail,
		Tel:  p.Tel,
	}
}
