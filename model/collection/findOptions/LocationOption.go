package findOptions

// 位置情報検索のオプション
type LocationOption interface {
	FindOption
	SetLocation(float64, float64, int)
	GetLocation() (float64, float64, float64)
	IsLocationSet() bool
}
