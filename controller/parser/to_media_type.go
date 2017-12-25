package parser

import (
	"EvelyApi/app"
	"EvelyApi/model"
)

// 緯度経度の配列番号を定数化
const (
	Lng = 0
	Lat = 1
)

/**
 * イベント情報をAPIのレスポンス形式に変換する
 * @param e イベント情報
 * @return  レスポンス形式に変換したイベント情報
 */
func ToEventMedia(e *model.EventModel) *app.Event {
	return &app.Event{
		ID:    e.ID,
		Title: e.Title,
		Body:  e.Body,
		Host: &app.UserTiny{
			ID:   e.Host.ID,
			Name: e.Host.Name,
		},
		Mail:        e.Mail,
		Tel:         e.Tel,
		URL:         e.URL,
		Plans:       toPlansMedia(e.Plans),
		NoticeRange: e.NoticeRange,
		Scope:       e.Scope,
		OpenFlg:     e.OpenFlg,
		UpdateDate:  e.UpdateDate,
		CreatedAt:   e.CreatedAt,
	}
}

/**
 * イベント情報をAPIのレスポンス形式に変換する
 * @param e イベント情報
 * @return  レスポンス形式に変換したイベント情報
 */
func ToEventTinyMedia(e *model.EventModel) *app.EventTiny {
	return &app.EventTiny{
		ID:    e.ID,
		Title: e.Title,
		Host: &app.UserTiny{
			ID:   e.Host.ID,
			Name: e.Host.Name,
		},
		Plans:       toPlansMedia(e.Plans),
		NoticeRange: e.NoticeRange,
		Scope:       e.Scope,
		OpenFlg:     e.OpenFlg,
		UpdateDate:  e.UpdateDate,
	}
}

func toPlansMedia(oldPlans []model.Plan) (newPlans []*app.Plan) {
	for _, old := range oldPlans {
		plan := &app.Plan{
			Location: &app.Location{
				Name: old.Location.Name,
				Lng:  old.Location.LngLat[Lng],
				Lat:  old.Location.LngLat[Lat],
			},
			UpcomingDate: &app.UpcomingDate{
				StartDate: old.UpcomingDate.StartDate,
				EndDate:   old.UpcomingDate.EndDate,
			},
		}
		newPlans = append(newPlans, plan)
	}
	return newPlans
}

/**
 * ユーザー情報をAPIのレスポンス形式に変換する
 * @param  u ユーザー情報
 * @return   レスポンス形式に変換したユーザー情報
 */
func ToUserMedia(u *model.UserModel) *app.User {
	return &app.User{
		ID:   u.ID,
		Name: u.Name,
		Mail: &app.Mail{
			Email: u.Mail.Email,
			State: u.Mail.State,
		},
		Tel:       u.Tel,
		CreatedAt: u.CreatedAt,
	}
}
