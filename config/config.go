package config

import (
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Station      Station
	Wunderground Wunderground
	MQTT         MQTT
}

type Station struct {
	ID                     string
	Password               string
	URL                    string
	WatchdogEnabled        bool
	RebootOnFailedAttempts int
	MQTTUpdateTopic        string
}

type Wunderground struct {
	Enabled   bool
	UpdateURL string
}

type MQTT struct {
	Enabled  bool
	Broker   string
	Username string
	Password string
	ClientID string
}

func NewConfig() *Config {
	c := NewConfigDefaults()
	if err := c.loadConfigFromFile("config.toml"); err != nil {
		log.Println("Unable to load config.toml, loaded defaults...")
	}
	c.applyEnvirontmentVariables()

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

func (c *Config) applyEnvirontmentVariables() {
	applyEnvirontmentVariable("STATION_ID", &c.Station.ID)
	applyEnvirontmentVariable("STATION_PASSWORD", &c.Station.Password)
	applyEnvirontmentVariable("STATION_URL", &c.Station.URL)
	applyEnvirontmentVariable("STATION_WATCHDOG_ENABLED", &c.Station.WatchdogEnabled)
	applyEnvirontmentVariable("STATION_REBOOT_ON_FAILED_ATTEMPTS", &c.Station.RebootOnFailedAttempts)
	applyEnvirontmentVariable("STATION_MQTT_UPDATE_TOPIC", &c.Station.MQTTUpdateTopic)

	applyEnvirontmentVariable("WUNDERGROUND_ENABLED", &c.Wunderground.Enabled)
	applyEnvirontmentVariable("WUNDERGROUND_UPDATE_URL", &c.Wunderground.UpdateURL)

	applyEnvirontmentVariable("MQTT_ENABLED", &c.MQTT.Enabled)
	applyEnvirontmentVariable("MQTT_BROKER", &c.MQTT.Broker)
	applyEnvirontmentVariable("MQTT_USERNAME", &c.MQTT.Username)
	applyEnvirontmentVariable("MQTT_PASSWORD", &c.MQTT.Password)
	applyEnvirontmentVariable("MQTT_CLIENT_ID", &c.MQTT.ClientID)
}

func applyEnvirontmentVariable(key string, value interface{}) {
	if env, ok := os.LookupEnv(key); ok {
		switch v := value.(type) {
		case *string:
			*v = env
		case *bool:
			if env == "true" || env == "1" {
				*v = true
			} else if env == "false" || env == "0" {
				*v = false
			}
		case *int:
			if number, err := strconv.Atoi(env); err == nil {
				*v = number
			}
		}
	}
}
