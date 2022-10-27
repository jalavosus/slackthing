package main

import (
	"github.com/jalavosus/slackthing/internal/config"
	"github.com/urfave/cli/v2"
)

var (
	configFileFlag = cli.PathFlag{
		Name:     "configFile",
		Usage:    "`path` to a configuration file",
		Aliases:  []string{"c"},
		Required: true,
	}
)

func loadConfigFile(c *cli.Context) (config.AppConfig, error) {
	configPath := configFileFlag.Get(c)
	return config.LoadConfig(configPath)
}
