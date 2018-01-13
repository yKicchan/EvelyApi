package mail

import (
    . "EvelyApi/config"
    "gopkg.in/gomail.v2"
)

/**
 * メールを送信する
 * @param  to      宛先
 * @param  subject 件名
 * @param  body    本文
 * @return error   メール送信時のエラー
 */
func sendMail(to, subject, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", "evely.ecc@gmail.com")
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer("smtp.gmail.com", 587, "evely.ecc@gmail.com", "123qwEcc")

    return d.DialAndSend(m)
}

func SendSignUpMail(to, url string) error {
    body := makeSignUpBody(to, url)
    return sendMail(to, SIGNUP_SUBJECT, body)
}
