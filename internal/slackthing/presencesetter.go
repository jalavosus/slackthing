package slackthing

import (
	"context"
	"github.com/jalavosus/slackthing/internal/config"
	"github.com/jalavosus/slackthing/internal/slackclient"
	"github.com/jalavosus/slackthing/internal/utils"
	"go.uber.org/zap"
	"time"
)

const (
	presenceSetterName           string = "PresenceSetter"
	defaultPresenceCheckInterval        = time.Minute
)

type PresenceSetter struct {
	*baseSlackThing
}

func NewPresenceSetter(cfg config.AppConfig) (SlackThing, error) {
	base, err := newBaseSlackThing(cfg, presenceSetterName)
	if err != nil {
		return nil, err
	}

	s := &PresenceSetter{
		baseSlackThing: base,
	}

	return s, nil
}

func (s *PresenceSetter) config() *config.PresenceSetterConfig {
	return s.cfg.PresenceSetter
}

func (s *PresenceSetter) Start(ctx context.Context) error {
	var (
		setActive     = utils.ParsedTime{Hour: -1}
		setAway       = utils.ParsedTime{Hour: -1}
		checkInterval = defaultPresenceCheckInterval
	)

	if cfg := s.config(); cfg != nil {
		var parseTimeErr error

		if activeTime := cfg.ActiveTime; activeTime != "" {
			setActive, parseTimeErr = utils.ParseTime(cfg.ActiveTime)
			if parseTimeErr != nil {
				return parseTimeErr
			}
		}

		if awayTime := cfg.AwayTime; awayTime != "" {
			setAway, parseTimeErr = utils.ParseTime(cfg.AwayTime)
			if parseTimeErr != nil {
				return parseTimeErr
			}
		}

		if configCheckInterval := cfg.CheckInterval; configCheckInterval != time.Duration(0) {
			checkInterval = configCheckInterval
		}
	}

	ticker := time.NewTicker(checkInterval)

	logStartFields := make([]zap.Field, 0, 3)

	logStartFields = append(logStartFields, zap.String("check_interval", checkInterval.String()))

	if setActive.Valid() {
		logStartFields = append(logStartFields, zap.String("active_time", setActive.Format()))
	}

	if setAway.Valid() {
		logStartFields = append(logStartFields, zap.String("away_time", setAway.Format()))
	}

	s.logStart(logStartFields...)

	for {
		select {
		case <-ticker.C:
			go s.doStuff(ctx, setActive, setAway)
		case <-ctx.Done():
			return nil
		}
	}
}

func (s *PresenceSetter) getPresence(ctx context.Context) (slackclient.UserPresence, error) {
	var presence slackclient.UserPresence

	_ctx, cancel := ctxFromCtx(ctx, clientTimeout)
	defer cancel()

	res, err := s.client.GetUserPresence(_ctx, s.cfg.UserId)
	if err != nil {
		return presence, err
	}

	presence = res

	return presence, nil
}

func (s *PresenceSetter) setPresence(ctx context.Context, presence slackclient.UserPresence) error {
	_ctx, cancel := ctxFromCtx(ctx, clientTimeout)
	defer cancel()

	return s.client.SetUserPresence(_ctx, presence)
}

func (s *PresenceSetter) doStuff(ctx context.Context, setActive, setAway utils.ParsedTime) {
	currentPresence, err := s.getPresence(ctx)
	if err != nil {
		s.logger.Error(
			"error fetching current presence",
			zap.Error(err),
		)
		return
	}

	var newPresence slackclient.UserPresence

	checkTime := utils.ToTimezone(time.Now())

	if isWeekend(checkTime) {
		newPresence = slackclient.Away
	} else {
		var canSet bool
		newPresence, canSet = getNewPresence(setActive, setAway, checkTime)
		if !canSet {
			return
		}
	}

	if checkEqualPresence(currentPresence, newPresence) {
		return
	}

	canSet := false
	if newPresence == slackclient.Active && setActive.Valid() {
		canSet = true
	} else if newPresence == slackclient.Away && setAway.Valid() {
		canSet = true
	}

	if canSet {
		if err = s.setPresence(ctx, newPresence); err != nil {
			s.logger.Error(
				"error setting presence",
				zap.Error(err),
			)
			return
		}

		s.logger.Info(
			"successfully set presence",
			zap.String("new_presence", newPresence.String()),
		)
	}

	return
}
