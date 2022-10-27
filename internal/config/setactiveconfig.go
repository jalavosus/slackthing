package config

import "time"

type PresenceSetterConfig struct {
	ActiveTime    string        `json:"active_time" yaml:"active_time"`
	AwayTime      string        `json:"away_time" yaml:"away_time"`
	CheckInterval time.Duration `json:"check_interval" yaml:"check_interval"`
}
