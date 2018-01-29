package parser

import (
	"EvelyApi/app"
	. "EvelyApi/models/documents"
	"labix.org/v2/mgo/bson"
)

/**
 * イベント情報をAPIのレスポンス形式に変換する
 * @param e イベント情報
 * @return  レスポンス形式に変換したイベント情報
 */
func ToEventMedia(e *EventModel) *app.Event {
	return &app.Event{
		ID:        e.ID.Hex(),
		Image:     toFileMedia(e.Image),
		Title:     e.Title,
		Body:      e.Body,
		Files:     toFilesMedia(e.Files),
		Categorys: e.Categorys,
		Host: &app.UserTiny{
			ID:   e.Host.ID,
			Name: e.Host.Name,
			Icon: toFileMedia(e.Host.Icon),
		},
		Mail:       e.Mail,
		Tel:        e.Tel,
		URL:        e.URL,
		Schedules:  toSchedulesMedia(e.Schedules),
		IsReviewed: len(e.Reviews) > 0,
		UpdateDate: e.UpdateDate,
		CreatedAt:  e.CreatedAt,
	}
}

/**
 * イベント情報をAPIのレスポンス形式に変換する
 * @param e イベント情報
 * @return  レスポンス形式に変換したイベント情報
 */
func ToEventTinyMedia(e *EventModel) *app.EventTiny {
	return &app.EventTiny{
		ID:        e.ID.Hex(),
		Image:     toFileMedia(e.Image),
		Title:     e.Title,
		Categorys: e.Categorys,
		Host: &app.UserTiny{
			ID:   e.Host.ID,
			Name: e.Host.Name,
			Icon: toFileMedia(e.Host.Icon),
		},
		Schedules:  toSchedulesMedia(e.Schedules),
		IsReviewed: len(e.Reviews) > 0,
	}
}

func toSchedulesMedia(oldSchedules []*Schedule) (newSchedules []*app.Schedule) {
	for _, old := range oldSchedules {
		schedule := &app.Schedule{
			Location: &app.Location{
				Name: old.Location.Name,
				Lng:  old.Location.LngLat[LNG],
				Lat:  old.Location.LngLat[LAT],
			},
			UpcomingDate: &app.UpcomingDate{
				StartDate: old.UpcomingDate.StartDate,
				EndDate:   old.UpcomingDate.EndDate,
			},
		}
		newSchedules = append(newSchedules, schedule)
	}
	return newSchedules
}

/**
 * ユーザー情報をAPIのレスポンス形式に変換する
 * @param  u ユーザー情報
 * @return   レスポンス形式に変換したユーザー情報
 */
func ToUserMedia(u *UserModel) *app.User {
	return &app.User{
		ID:   u.ID,
		Name: u.Name,
		Icon: toFileMedia(u.Icon),
		Mail: &app.Mail{
			Email: u.Mail.Email,
			State: u.Mail.State,
		},
		Tel:         u.Tel,
		Pins:        toPinsMedia(u.Pins),
		Preferences: u.Preferences,
		CreatedAt:   u.CreatedAt,
	}
}

func toPinsMedia(old []bson.ObjectId) (new []string) {
	for _, o := range old {
		new = append(new, o.Hex())
	}
	return new
}

/**
 * メールアドレスをレスポンス形式に変換する
 * @param  email メールアドレス
 * @return Email レスポンス形式に変換したメールアドレス
 */
func ToEmailMedia(email string) *app.Email {
	return &app.Email{
		Email: email,
	}
}

/**
 * レビューをレスポンス形式に変換する
 * @param  r      レビュー
 * @return Review レスポンス形式に変換したレビュー
 */
func ToReviewMedia(r *ReviewModel) *app.Review {
	return &app.Review{
		ID:    r.ID.Hex(),
		Rate:  r.Rate,
		Title: r.Title,
		Body:  r.Body,
		Files: toFilesMedia(r.Files),
		Reviewer: &app.UserTiny{
			ID:   r.Reviewer.ID,
			Name: r.Reviewer.Name,
			Icon: toFileMedia(r.Reviewer.Icon),
		},
		ReviewedAt: r.ReviewedAt,
	}
}

func toFileMedia(f string) string {
	return "http://localhost:8888/download/" + f
}

func toFilesMedia(old []string) (new []string) {
	for _, o := range old {
		new = append(new, toFileMedia(o))
	}
	return new
}
