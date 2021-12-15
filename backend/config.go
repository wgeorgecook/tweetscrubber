package main

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

type config struct {
	TwitterApiKey      string `env:"TWITTER_API_KEY"`
	TwitterApiSecret   int    `env:"TWITTER_API_SECRET"`
	TwitterBearerToken string `env:"TWITTER_BEARER_TOKEN"`
}

var logger *zap.SugaredLogger

func initLogging() {
	logConfig, _ := zap.NewProduction()
	defer logConfig.Sync() // flushes buffer, if any
	logger = logConfig.Sugar()
}

func initEnv() error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	logger.Info("config loaded")
	return nil
}
