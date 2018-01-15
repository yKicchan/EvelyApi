package mailer

const (

	// 登録確認メール内容
	SIGNUP_CONFIRM_SUBJECT      = "メールアドレスの確認"
	SIGNUP_CONFIRM_BODY_TITLE   = "メールアドレスを<br>認証してください"
	SIGNUP_CONFIRM_BODY_MESSAGE = "ユーザー登録はまだ完了していません<br>下記のボタンからメールアドレスを認証し、登録を完了してください"

	// 登録完了メールの内容
	SIGNUP_COMPLETE_SUBJECT = "ユーザー登録が完了しました"
)

/**
 * 登録確認メールの本文を返す
 * @param  to     送信先メールアドレス
 * @param  url    登録用URL
 * @return string 登録確認メール本文
 */
func makeSignUpConfirmBody(to, url string) string {
	return `
<html>
<head>
<title>` + SIGNUP_CONFIRM_BODY_TITLE + `</title>
</head>
<body>
<table width="100%" height="100%" style="min-width:348px;" border="0" cellspacing="0" cellpadding="0">
    <tbody>
        <tr align="center">
            <td>
                <table border="0" cellspacing="0" cellpadding="0" style="padding-bottom:20px;max-width:500px;min-width:220px;border-top:4px solid #ffcd8e;border-right:1px solid #dcdcdc;border-bottom:1px solid #dcdcdc;border-left:1px solid #dcdcdc;border-radius:4px">
                    <tr>
                        <td>
                            <div style="font-family:Arial,sans-serif;padding:30px 20px 10px 20px;color:rgba(0,0,0,0.87);font-size:20px;text-align:center;">
                                <div>
                                    ` + SIGNUP_CONFIRM_BODY_TITLE + `<br>
                                    <span style="color:rgba(0,0,0,0.87);font-size:0.85em;text-decoration:none">` + to + `</span>
                                </div>
                            </div>
                            <div style="border-top:1px solid #e8e8e8;width:100%;"></div>
                            <div style="font-family:Arial,sans-serif;font-size:0.8em;color:rgba(0,0,0,0.87);line-height:1.6em;padding:30px 20px 0px;">
                                <div>
                                    ` + SIGNUP_CONFIRM_BODY_MESSAGE + `
                                    <div style="padding-top:24px;text-align:center">
                                        <a href="` + url + `" style="display:inline-block;text-decoration:none;" target="_blank">
                                            <table border="0" style="background-color:#4184f3;border-radius:3px;width:80px">
                                                <tr>
                                                    <td style="padding-left:7px;padding-right:7px;text-align:center">
                                                        <a href="` + url + `" style="font-family:Arial,sans-serif;color:#ffffff;text-decoration:none;font-size:0.85em;line-height:1.8em;">認証する</a>
                                                    </td>
                                                </tr>
                                            </table>
                                        </a>
                                    </div>
                                </div>
                            </div>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </tbody>
</table>
</body>
</html>
`
}

/**
 * 登録完了メールの本文を返す
 * @param  name   送信先ユーザーのニックネーム
 * @return string 登録完了メール本文
 */
func makeSignUpCompleteBody(name string) string {
	return name + `さん Evelyへようこそ！

Evelyでは近くのイベントや、興味のあるイベントを探すことができます。
チョットお得な情報や、流行りのイベントを探し出して、毎日を楽しく過ごしましょう！

Evelyのイベント検索・通知の受け取りにはスマホアプリがオススメです。
☆GooglePlayStore
[URL]

イベントの作成・管理をPCから行いたいという方は、Webブラウザからもお使いいただけます。
☆Evely公式URL
[URL]

☆お問い合わせはこちら
[URL]

------------------------------------------------
Evely運営チーム: Opportunes
責任元: ECCコンピュータ専門学校
〒530-0015 大阪府大阪市 北区中崎西2丁目3−35
------------------------------------------------
`
}
