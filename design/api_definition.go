package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// APIの定義
var _ = API("EvelyApi", func() {

	// APIの説明
	Title("EvelyApi")
	Description("卒業制作のアプリケーションで使用するAPIです")

	// 作成者へのコンタクト
	Contact(func() {
		Name("yKicchan")
		Email("yKicchanApp@gmail.com")
	})

	// ライセンス情報
	License(func() {
		Name("Apache License 2.0")
		URL("https://github.com/yKicchan/EvelyApi/blob/master/LICENSE")
	})

	// ホスト設定
	Scheme("http")
	Host("localhost:8888")

	// 全エンドポイントのBasePath
	BasePath("/api/develop/v1")

	// CORSポリシーの定義
	Origin("http://localhost:8888/swaggerui", func() {
		// クライアントに公開された1つ以上のヘッダー
		Expose("X-Time")
		// 許可されたHTTPメソッド
		Methods("GET", "POST", "PUT", "DELETE")
		// プリフライト要求応答をキャッシュする時間
		MaxAge(600)
		// Access-Control-Allow-Credentialsヘッダーを設定する
		Credentials()
	})

	// リソースに対するアクセス権限
	Security(JWT, func() {
		Scope("api:access")
	})
})
