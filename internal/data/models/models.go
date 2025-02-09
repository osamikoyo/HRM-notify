package models

import "gopkg.in/gomail.v2"

type Notify struct{
	ReciewerEmail string
	Subject string
	Content string
}

type Message struct{
	*gomail.Message
}

func NewMessage(from, to string, addressHeader []string, subject string, body string) *Message {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", addressHeader[0], addressHeader[1])
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return &Message{m}
}