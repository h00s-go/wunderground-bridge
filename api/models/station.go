package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/wunderground-bridge/config"
	"github.com/h00s-go/wunderground-bridge/mqtt"
)

type Station struct {
	config   *config.Station
	logger   *log.Logger
	Watchdog Watchdog
	Weather  *Weather
}

type Watchdog struct {
	SuccessfulLastUpdate bool
	FailedAttempts       int
}

func NewStation(config *config.Station, logger *log.Logger) *Station {
	return &Station{
		config:   config,
		logger:   logger,
		Watchdog: Watchdog{},
		Weather:  &Weather{},
	}
}

func (s *Station) UpdateWeather(ctx *fiber.Ctx) error {
	if ctx.Query("ID") != s.config.ID {
		return errors.New("Station ID does not match")
	}
	if ctx.Query("PASSWORD") != s.config.Password {
		return errors.New("Station password does not match")
	}

	w, err := NewWeather(ctx)
	s.UpdateWatchDog(err)
	if err != nil {
		return err
	}
	s.Weather = w
	return nil
}

func (s *Station) UpdateWatchDog(updateError error) {
	if s.config.WatchdogEnabled {
		if updateError == nil {
			s.Watchdog.SuccessfulLastUpdate = true
			s.Watchdog.FailedAttempts = 0
			return
		}
		if s.Watchdog.FailedAttempts >= s.config.RebootOnFailedAttempts {
			s.logger.Println("Attempting reboot")
			go s.attemptReboot()
			s.Watchdog.FailedAttempts = 0
		}
		s.Watchdog.SuccessfulLastUpdate = false
		s.Watchdog.FailedAttempts++
		s.logger.Printf("Failed %v time(s) to update weather data", s.Watchdog.FailedAttempts)
	}
}

func (s *Station) attemptReboot() {
	_, err := http.Get(fmt.Sprintf("%s/msgreboot.htm", s.config.URL))
	if err != nil {
		s.logger.Printf("Error while reboot attempt: %v\n", err)
	}
}

func (s *Station) PublishWeatherToMQTT(mqtt *mqtt.MQTT) {
	if mqtt.Config.Enabled {
		w, err := json.Marshal(s.Weather)
		if err == nil {
			mqtt.Publish(s.config.MQTTUpdateTopic, string(w))
		}
	}
}
