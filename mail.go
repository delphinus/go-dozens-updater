package godo

import (
	"bytes"
	"net/smtp"
	"os"
	"strings"
)

var auth = smtp.PlainAuth(
	"",
	os.Getenv("GODO_MAIL_USER"),
	os.Getenv("GODO_MAIL_PASSWORD"),
	"smtp.gmail.com",
)

func sendMail(content []byte) error {
	recipients := strings.Split(os.Getenv("GODO_MAIL_RECIPIENTS"), ",")
	buf := bytes.NewBufferString("Subject: ")
	_, _ = buf.WriteString(os.Getenv("GODO_MAIL_SUBJECT"))
	_, _ = buf.WriteString("\r\n\r\n")
	_, _ = buf.Write(content)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("GODO_MAIL"),
		recipients,
		buf.Bytes(),
	)
}
