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
		Body:  p.Body,
		Host: model.Host{
			ID:   user.ID,
			Name: user.Name,
		},
		Mail:        p.Mail,
		Tel:         p.Tel,
		URL:         p.URL,
		Plans:       toPlansModel(p.Plans),
		NoticeRange: p.NoticeRange,
		Scope:       p.Scope,
		OpenFlg:     p.OpenFlg,
		UpdateDate:  time.Now(),
	}
}

func toPlansModel(oldPlans []*app.Plan) (newPlans []model.Plan) {
	for _, old := range oldPlans {
		plan := model.Plan{
			Location: model.Location{
				Name:   old.Location.Name,
				LngLat: [2]float64{old.Location.Lng, old.Location.Lat},
			},
			UpcomingDate: model.UpcomingDate{
				StartDate: old.UpcomingDate.StartDate,
				EndDate:   old.UpcomingDate.EndDate,
			},
		}
		newPlans = append(newPlans, plan)
	}
	return newPlans
}
