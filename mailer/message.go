package mailer

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
	"jaytaylor.com/html2text"
)

type Message struct {
	Info    string
	To      string
	Subject string
	Body    string
	Headers map[string][]string
}

func (m *Mailer) SendMessage(msg *Message) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		defer func() {
			if r := recover(); r != nil {
				m.logger.Err(fmt.Errorf("panic while sending message: %v", r)).Str("info", msg.Info).Msg("failed to send email")
			}
		}()

		err := m.SendMessageSync(msg)

		if err != nil {
			m.logger.Err(err).Str("info", msg.Info).Msg("failed to send email")
			return
		}
	}()
}

func (m *Mailer) SendMessageSync(msg *Message) error {
	mailMsg := gomail.NewMessage()

	mailMsg.SetAddressHeader("From", m.config.FromAddress, m.config.FromDisplayName)
	mailMsg.SetHeader("To", msg.To)
	if m.config.ReplyTo != "" {
		mailMsg.SetHeader("Reply-To", m.config.ReplyTo)
	}
	for header := range msg.Headers {
		mailMsg.SetHeader(header, msg.Headers[header]...)
	}
	mailMsg.SetHeader("Subject", msg.Subject)
	mailMsg.SetDateHeader("Date", time.Now())

	plainBody, err := html2text.FromString(msg.Body, html2text.Options{
		PrettyTables: true,
	})
	if err != nil {
		return err
	}

	mailMsg.SetBody("text/plain", plainBody)
	mailMsg.AddAlternative("text/html", msg.Body)
	return m.dialer.DialAndSend(mailMsg)
}
