package findOptions

// 位置情報検索のオプション
type LocationOption interface {
	FindOption
	SetLocation(float64, float64, int)
	GetLocation() (float64, float64, int)
	IsLocationSet() bool
}

type locationOption struct {
	// 緯度(0があり得るのでセット判定のためにポインタ型)
	lat *float64
	// 経度(0があり得るのでセット判定のためにポインタ型)
	lng *float64
	// 範囲(単位:度)
	r int
}

/**
 * 設定された位置情報と範囲を返す
 * @return lat 緯度
 * @return lng 経度
 * @return r   範囲(度)
 */
func (this *locationOption) GetLocation() (lat, lng float64, r int) {
	return *this.lat, *this.lng, this.r
}

/**
 * 位置情報を検索オプションに設定する
 * @param  lat 緯度
 * @param  lng 経度
 * @param  r   検索範囲(m)
 */
func (this *locationOption) SetLocation(lat, lng float64, r int) {
	if lat <= 90 && lat >= -90 {
		this.lat = &lat
	}
	if lng <= 180 && lng >= -180 {
		this.lng = &lng
	}
	if r > 0 {
		this.r = r
	}
}

/**
 * 位置情報検索が有効かを判定する
 * @return bool
 */
func (this *locationOption) IsLocationSet() bool {
	return this.lat != nil && this.lng != nil && this.r > 0.0
}
