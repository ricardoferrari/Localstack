package bootstrap

import (
	"context"
	"log"
	"os"

	"go.uber.org/fx"
)

func NewLogger(lc fx.Lifecycle) *log.Logger {
	logger := log.New(os.Stdout, "[GOREST] " /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			logger.Print("Starting logger.")
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Print("Shutting down logger.")
			return nil
		},
	})
	return logger
}
