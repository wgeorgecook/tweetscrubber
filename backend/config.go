package main

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

type cfg struct {
	TwitterApiKey      string `env:"TWITTER_API_KEY"`
	TwitterApiSecret   int    `env:"TWITTER_API_SECRET"`
	TwitterBearerToken string `env:"TWITTER_BEARER_TOKEN"`
	TwitterClientID    string `env:"TWITTER_CLIENT_ID"`
	TwitterRedirectURI string `env:"TWITTER_REDIRECT_URI`
}

var logger *zap.SugaredLogger
var config cfg

func initLogging() {
	logConfig, _ := zap.NewProduction()
	defer logConfig.Sync() // flushes buffer, if any
	logger = logConfig.Sugar()
}

func initEnv() error {
	config = cfg{}
	if err := env.Parse(&config); err != nil {
		return err
	}

	logger.Info("config loaded")
	return nil
}
