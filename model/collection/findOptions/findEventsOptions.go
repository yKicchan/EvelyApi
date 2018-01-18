package findOptions

import (
	. "EvelyApi/config"
)

// イベント検索で使用する検索オプション
type findEventsOptions struct {
	findOptions
	// キーワード
	keyword string
	// 緯度(0があり得るのでセット判定のためにポインタ型)
	lat *float64
	// 経度(0があり得るのでセット判定のためにポインタ型)
	lng *float64
	// 範囲(単位:度)
	r float64
}

func NewFindEventsOptions() *findEventsOptions {
	return &findEventsOptions{}
}

/**
 * キーワードを検索オプションに設定する
 * @param  keyword キーワード
 */
func (this *findEventsOptions) SetKeyword(keyword string) {
	if keyword != "" {
		this.keyword = keyword
	}
}

/**
 * 設定されたキーワードを返す
 * @return キーワード
 */
func (this *findEventsOptions) GetKeyword() string { return this.keyword }

/**
 * キーワードが有効かを判定する
 * @return bool
 */
func (this *findEventsOptions) IsKeywordSet() bool { return this.keyword != "" }

/**
 * 設定された位置情報と範囲を返す
 * @return lat 緯度
 * @return lng 経度
 * @return r   範囲(度)
 */
func (this *findEventsOptions) GetLocation() (lat, lng, r float64) {
	return *this.lat, *this.lng, this.r
}

/**
 * 位置情報を検索オプションに設定する
 * @param  lat 緯度
 * @param  lng 経度
 * @param  r   検索範囲(m)
 */
func (this *findEventsOptions) SetLocation(lat, lng float64, r int) {
	if lat <= 90 && lat >= -90 {
		this.lat = &lat
	}
	if lng <= 180 && lng >= -180 {
		this.lng = &lng
	}
	if r > 0 {
		this.r = float64(r) * DEGREE_PER_METER
	}
}

/**
 * 位置情報検索が有効かを判定する
 * @return bool
 */
func (this *findEventsOptions) IsLocationSet() bool {
	return this.lat != nil && this.lng != nil && this.r > 0.0
}
