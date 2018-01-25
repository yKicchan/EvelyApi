package mailer

import (
	"gopkg.in/gomail.v2"
)

/**
 * メールを送信する
 * @param  to      宛先
 * @param  subject 件名
 * @param  body    本文
 * @param  html    true: HTMLメール, false: 普通のテキストメール
 * @return error   メール送信時のエラー
 */
func sendMail(to, subject, body string, html bool) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "Evely Operation Team<evely.ecc@gmail.com>")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "[Evely]"+subject)
	if html {
		m.SetBody("text/html", body)
	} else {
		m.SetBody("text/plain", body)
	}

	d := gomail.NewDialer("smtp.gmail.com", 587, "evely.ecc@gmail.com", "123qwEcc")

	return d.DialAndSend(m)
}

/**
 * 登録確認メールを送信する
 * @param to 送信先メールアドレス
 * @param url メールに記載する登録用URL
 * @return error メール送信時に発生したエラー内容
 */
func SendSignUpConfirmMail(to, url string) error {
	body := makeSignUpConfirmBody(to, url)
	return sendMail(to, SIGNUP_CONFIRM_SUBJECT, body, true)
}

/**
 * 登録完了メールを送信する
 * @param to 送信先メールアドレス
 * @param name 送信先ユーザーのニックネーム
 * @return error メール送信時に発生したエラー
 */
func SendSignUpCompleteMail(to, name string) error {
	body := makeSignUpCompleteBody(name)
	return sendMail(to, SIGNUP_COMPLETE_SUBJECT, body, false)
}
