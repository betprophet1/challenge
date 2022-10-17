package app

import (
	"os"
	"project/config"
	"time"
)

type Options struct {
	GracefulTimeout time.Duration
}

func parseOpt() Options {
	gracefulTimeout, _ := time.ParseDuration(os.Getenv(config.SERVICE_GRACEFUL_TIMEOUT))
	return Options{
		GracefulTimeout: gracefulTimeout,
	}
}
