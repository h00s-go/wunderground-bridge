package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Station      Station
	Wunderground Wunderground
	MQTT         MQTT
}

type Station struct {
	URL                    string
	WatchdogEnabled        bool
	RebootOnFailedAttempts int
}

type Wunderground struct {
	Enabled   bool
	UpdateURL string
}

type MQTT struct {
	Enabled     bool
	Broker      string
	Port        string
	Username    string
	Password    string
	ClientID    string
	UpdateTopic string
}

func NewConfig() *Config {
	c := NewConfigDefaults()
	if err := c.loadConfigFromFile("config.toml"); err != nil {
		log.Println("Unable to load config.toml, loaded defaults...")
	}

	return c
}

func NewConfigDefaults() *Config {
	return &Config{
		Station: Station{
			RebootOnFailedAttempts: 15,
		},
		Wunderground: Wunderground{
			Enabled:   true,
			UpdateURL: "http://weatherstation.wunderground.com/weatherstation/updateweatherstation.php",
		},
	}
}

func (c *Config) loadConfigFromFile(path string) error {
	if _, err := toml.DecodeFile(path, c); err != nil {
		return err
	}

	return nil
}
