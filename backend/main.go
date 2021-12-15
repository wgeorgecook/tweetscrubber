package main

func init() {
	initLogging()
	if err := initEnv(); err != nil {
		logger.Infof("could not read .env: %v", err)
	}
}

func main() {
	logger.Info("Start main")
	startHttpServer()
	defer logger.Info("Goodbye!")
}
