package main

import (
	"net/smtp"

	"github.com/aitorfernandez/roshambo/pkg/env"
)

// sender sends emails.
type sender struct {
	addr, host, identity, password, username string
}

// DefaultSender is the default Sender.
var DefaultSender = &sender{
	addr:     env.MustHget("gmail", "addr"),
	host:     env.MustHget("gmail", "host"),
	identity: env.MustHget("gmail", "identity"),
	password: env.MustHget("gmail", "password"),
	username: env.MustHget("gmail", "username"),
}

func (s sender) send(to []string, subject, body string) error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	plainAuth := smtp.PlainAuth(s.identity, s.username, s.password, s.host)
	msg := []byte("Subject: " + subject + "\n" + mime + body)
	return smtp.SendMail(s.addr, plainAuth, s.username, to, msg)
}

func send(to []string, subject, body string) error {
	return DefaultSender.send(to, subject, body)
}
