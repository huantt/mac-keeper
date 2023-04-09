package alerter

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
)

type SlackService struct {
	webhookUrl string
	channel    string
}

func NewSlackService(webhookUrl string, channel string) *SlackService {
	return &SlackService{
		webhookUrl: webhookUrl,
		channel:    channel,
	}
}

func (s *SlackService) Alert(message string) error {
	payload := slack.Payload{
		Text:    fmt.Sprintf("[House keeper] %s", message),
		Channel: s.channel,
	}
	err := slack.Send(s.webhookUrl, "", payload)
	if len(err) > 0 {
		return err[0]
	}
	return nil
}
