package config

const (

	// DB設定
	DB_HOST = "mongo"
	DB_NAME = "develop"

	// コレクション名
	EVENT_COLLECTION        = "Events"
	USER_COLLECTION         = "Users"
	PENDING_USER_COLLECTION = "PendingUsers"

	// メールアドレスの状態
	STATE_PENDING = "Pending"
	STATE_OK      = "OK"
	STATE_BAN     = "BAN"
)