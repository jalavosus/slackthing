package main

import (
	"context"
	"github.com/jalavosus/slackthing/internal/logging"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

var logger = logging.NewLoggerFromEnv("cmd")

func main() {
	app := &cli.App{
		Name:  "slackthing",
		Usage: "Set slack status automatically",
		Commands: []*cli.Command{
			&presenceSetterCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		if !errors.Is(err, context.Canceled) {
			logger.Fatal("application error", zap.Error(err))
		}

		return
	}
}
