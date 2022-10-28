package slackclient

import (
	"context"

	"github.com/slack-go/slack"
)

func (s *SlackClient) GetUserPresence(ctx context.Context, userId string) (UserPresence, error) {
	var (
		err      error
		presence = Unknown
		res      *slack.UserPresence
	)

	res, err = s.client.GetUserPresenceContext(ctx, userId)
	if err != nil {
		return presence, err
	}

	switch res.Presence {
	case "active":
		presence = Active
	case "auto":
		presence = Auto
	case "away":
		presence = Away
	}

	return presence, nil
}

func (s *SlackClient) SetUserPresence(ctx context.Context, presence UserPresence) error {
	if err := s.client.SetUserPresenceContext(ctx, presence.String()); err != nil {
		return err
	}

	return nil
}
