package slackclient

import (
	"context"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"time"
)

type SlackClient struct {
	client *slack.Client
}

func NewSlackClient(token, userId string) (*SlackClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := slack.New(token)
	if _, err := client.GetUserInfoContext(ctx, userId); err != nil {
		return nil, errors.WithMessage(err, "error initializing Slack client")
	}

	s := &SlackClient{client: client}

	return s, nil
}
