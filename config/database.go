package config

const (

	// DB設定
	DB_HOST = "mongo"
	DB_NAME = "develop"

	// コレクション名
	EVENTS_COLLECTION  = "Events"
	USERS_COLLECTION   = "Users"
    REVIEWS_COLLECTION = "Reviews"

	// メールアドレスの状態
	STATE_PENDING = "Pending"
	STATE_OK      = "OK"
	STATE_BAN     = "BAN"
    STATE_GUEST   = "Guest"
)
