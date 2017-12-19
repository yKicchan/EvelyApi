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
		Host: &app.UserTiny{
			ID:   e.Host.ID,
			Name: e.Host.Name,
		},
		Body: e.Body,
		Place: &app.Location{
			Name: e.Place.Name,
			Lat:  e.Place.Lng_Lat[Lat],
			Lng:  e.Place.Lng_Lat[Lng],
		},
		UpdateDate: e.Update_Date,
		UpcomingDate: &app.UpcomingDate{
			StartDate: e.Upcoming_Date.Start_Date,
			EndDate:   e.Upcoming_Date.End_Date,
		},
		URL:  e.URL,
		Mail: e.Mail,
		Tel:  e.Tel,
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
		Place: &app.Location{
			Name: e.Place.Name,
			Lat:  e.Place.Lng_Lat[Lat],
			Lng:  e.Place.Lng_Lat[Lng],
		},
		UpcomingDate: &app.UpcomingDate{
			StartDate: e.Upcoming_Date.Start_Date,
			EndDate:   e.Upcoming_Date.End_Date,
		},
	}
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
		Tel: u.Tel,
	}
}