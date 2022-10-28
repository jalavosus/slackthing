package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/jalavosus/slackthing/internal/slackthing"
	"github.com/jalavosus/slackthing/internal/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var presenceSetterCmd = cli.Command{
	Name:  "presence-setter",
	Usage: "StartProcess PresenceSetter slackthing",
	Flags: []cli.Flag{
		&configFileFlag,
	},
	Action: startPresenceSetter,
}

func startPresenceSetter(c *cli.Context) error {
	cfg, err := loadConfigFile(c)
	if err != nil {
		return err
	}

	thing, err := slackthing.NewPresenceSetter(cfg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(c.Context)
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	errCh := utils.StartProcess(ctx, thing)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-sigChan:
			logger.Warn(
				"received signal from os",
				zap.String("signal", sig.String()),
			)

			return nil
		case err = <-errCh:
			return err
		}
	}
}
