package config

import (
	"encoding/json"
	"github.com/jalavosus/slackthing/internal/utils"
	"github.com/pkg/errors"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type unmarshalFunc func([]byte, any) error

type AppConfig struct {
	UserId         string                `json:"user_id" yaml:"user_id"`
	SlackConfig    SlackConfig           `json:"slack" yaml:"slack"`
	PresenceSetter *PresenceSetterConfig `json:"presence_setter,omitempty" yaml:"presence_setter,omitempty"`
}

func validateAppConfig(c AppConfig) error {
	if c.SlackConfig.OauthToken == "" {
		return NewValidationError("AppConfig.SlackConfig.OauthToken", "cannot be empty")
	}

	if c.UserId == "" {
		return NewValidationError("AppConfig.UserId", "cannot be empty")
	}

	return nil
}

func LoadConfig(configPath string) (appConfig AppConfig, err error) {
	var cfgPath string
	appConfig = AppConfig{}

	cfgPath, err = utils.AbsoluteFilePath(configPath)
	if err != nil {
		err = errors.WithMessagef(err, "error parsing file psth %s", configPath)
		return
	}

	unmarshal := getUnmarshaler(cfgPath)

	var raw []byte
	raw, err = utils.ReadFile(cfgPath)
	if err != nil {
		err = errors.WithMessagef(err, "error reading from file %s", cfgPath)
		return
	}

	err = unmarshal(raw, &appConfig)
	if err != nil {
		err = errors.WithMessage(err, "error unmarshalling data")
		return
	}

	err = validateAppConfig(appConfig)
	if err != nil {
		err = errors.WithMessage(err, "validation error")
		return
	}

	return
}

func getUnmarshaler(configPath string) (fn unmarshalFunc) {
	switch filepath.Ext(configPath) {
	case ".yaml", ".yml":
		fn = yaml.Unmarshal
	case ".json":
		fn = json.Unmarshal
	}

	return
}
