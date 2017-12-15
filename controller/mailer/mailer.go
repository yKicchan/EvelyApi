package mailer

import (
    "net/smtp"
)

func SendMail(to, subject, body string) error {
        // Set up authentication information.
        auth := smtp.PlainAuth(
            "",
            "ykicchanapp@gmail.com",
            "Rathalos2",
            "smtp.gmail.com",
        )
        // Connect to the server, authenticate, set the sender and recipient,
        // and send the email all in one step.
        err := smtp.SendMail(
            "smtp.gmail.com:587",
            auth,
            "ykicchanapp@gmail.com", //foo@gmail.com
            []string{to},
            []byte(
                "To: " + to + "\r\n" +
                "Subject:[Evely] " + subject + "\r\n" +
                "\r\n" +
                body,
            ),
        )
        return err
}
