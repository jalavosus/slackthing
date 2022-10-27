package slackthing

import (
	"context"
	"github.com/jalavosus/slackthing/internal/config"
	"github.com/jalavosus/slackthing/internal/logging"
	"github.com/jalavosus/slackthing/internal/slackclient"
	"go.uber.org/zap"
	"time"
)

const (
	clientTimeout = 15 * time.Second
)

type SlackThing interface {
	Start(ctx context.Context) error
	Name() string
	logStart()
}

type baseSlackThing struct {
	cfg    config.AppConfig
	name   string
	client *slackclient.SlackClient
	logger *zap.Logger
}

func newBaseSlackThing(cfg config.AppConfig, name string) (*baseSlackThing, error) {
	client, err := slackclient.NewSlackClient(cfg.SlackConfig.OauthToken, cfg.UserId)
	if err != nil {
		return nil, err
	}

	s := &baseSlackThing{
		cfg:    cfg,
		name:   name,
		client: client,
		logger: logging.NewLoggerFromEnv(name),
	}

	return s, nil
}

func (s *baseSlackThing) Start(context.Context) error {
	panic("not implemented by baseSlackThing")
}

func (s *baseSlackThing) Name() string {
	return s.name
}

func (s *baseSlackThing) logStart() {
	s.logger.Info("starting SlackThing")
}
