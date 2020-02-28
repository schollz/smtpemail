package smtpemail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"net/textproto"
	"os"

	"github.com/jordan-wright/email"
	"github.com/yuin/goldmark"
)

const SMTPHost = "smtp-relay.sendinblue.com"
const SMTPPort = "465"

func Send(to, from, subject, markdown, attachment string) (err error) {
	smtpAuth := os.Getenv("SMTPAUTH")
	smtpPass := os.Getenv("SMTPPASS")
	if smtpAuth == "" || smtpPass == "" {
		err = fmt.Errorf("Must define environmental variables SMTPAUTH and SMTPPASS")
	}

	var buf bytes.Buffer
	if err = goldmark.Convert([]byte(markdown), &buf); err != nil {
		return
	}

	e := &email.Email{
		To:      []string{to},
		From:    from,
		Subject: subject,
		Text:    []byte(markdown),
		HTML:    buf.Bytes(),
		Headers: textproto.MIMEHeader{},
	}
	if attachment != "" {
		e.AttachFile(attachment)
	}
	err = e.SendWithTLS(SMTPHost+":"+SMTPPort, smtp.PlainAuth("", smtpAuth, smtpPass, SMTPHost), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTPHost,
	})
	return
}
