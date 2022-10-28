package slackthing

import (
	"context"
	"time"

	"github.com/jalavosus/slackthing/internal/config"
	"github.com/jalavosus/slackthing/internal/logging"
	"github.com/jalavosus/slackthing/internal/slackclient"
	"github.com/jalavosus/slackthing/internal/utils"
	"go.uber.org/zap"
)

const (
	clientTimeout = 15 * time.Second
)

type SlackThing interface {
	Start(ctx context.Context) error
	Name() string
	logStart(...zap.Field)
}

type baseSlackThing struct {
	cfg    config.AppConfig
	name   string
	client *slackclient.SlackClient
	logger *zap.Logger
}

func newBaseSlackThing(cfg config.AppConfig, name string) (*baseSlackThing, error) {
	client, err := slackclient.NewSlackClient(utils.SlackToken(), cfg.UserId)
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

func (s *baseSlackThing) logStart(fields ...zap.Field) {
	s.logger.With(fields...).Info("starting SlackThing")
}
