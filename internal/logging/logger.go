package logging

import "go.uber.org/zap"

func NewLogger(name string) *zap.Logger {
	l, _ := zap.NewProduction(
		zap.Fields(
			zap.String("name", name),
		),
	)

	return l
}
