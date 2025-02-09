package sender

import (
	"github.com/osamikoyo/hrm-notify/internal/data/models"
	"github.com/osamikoyo/hrm-notify/pkg/config"
	"gopkg.in/gomail.v2"
)

type Sender struct{
	*gomail.Dialer
}

func New(cfg *config.Config) *Sender {
	return &Sender{
		gomail.NewDialer(cfg.Smpt.SmptHost, cfg.Smpt.SmptPort, cfg.Smpt.SmptUsername, cfg.Smpt.SmptPassword),
	}
}

func (s *Sender) Send(message *models.Message) error {
	return s.DialAndSend(message.Message)
}