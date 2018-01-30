package config

// Evelyの独自定数
const (

	// 最大通知範囲
	MAX_NOTICE_RANGE = 5000

    // １メートルあたりの度(緯度基準)
    DEGREE_PER_METER = 0.000008983148616

	// Firebaseのサーバーキー
	FCM_SERVER_KEY = "AAAAW_wL-Tg:APA91bFOZ1pfJ_LFUiNaqFpaN6nlAkHHYHVmMNntKLWuTxaBDrQpBtuskKvvQgNVADvqNX-zP-S8qV4bs_oFkDyk0-a8WH_yAa5QqTcl5_KvLzADUHhoOHIPOnUWk-NbSFyVkQillUHe"

	// Evelyで独自定義したカテゴリ
	C_MUSIC = "Music"
	C_SPORTS = "Sports"
	C_VOLUNTEER = "Volunteer"
	C_ENTERTAINMENTS = "Entertainments"
	C_WORK_CONF = "Workshop & Conference"
	C_FOOD_DRINK = "Food & Drink"
	C_ARTS = "Arts"
	C_FESTIVAL = "Festival"
	C_BARGAIN = "Bargain"
)
// カテゴリ配列
type categorys []string
var Categorys = categorys{C_ARTS, C_BARGAIN, C_ENTERTAINMENTS, C_FESTIVAL, C_FOOD_DRINK, C_MUSIC, C_SPORTS, C_VOLUNTEER, C_WORK_CONF}

func (this categorys) IndexOf(value string) int {
    for i, v := range this {
        if (v == value) {
            return i
        }
    }
    return -1
}
