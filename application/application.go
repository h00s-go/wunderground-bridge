package application

import (
	"log"

	"github.com/h00s-go/wunderground-bridge/config"
)

type Application struct {
	config *config.Config
	logger *log.Logger
}

func NewApplication(config *config.Config, logger *log.Logger) *Application {
	return &Application{
		config: config,
		logger: logger,
	}
}
