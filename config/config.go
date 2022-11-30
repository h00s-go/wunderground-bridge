package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Wunderground Wunderground
	MQTT         MQTT
}

type Wunderground struct {
	PassUpdate bool
	UpdateURL  string
}

type MQTT struct {
	Broker       string
	Port         string
	Username     string
	Password     string
	ClientID     string
	WeatherTopic string
}

func NewConfig(path string) (*Config, error) {
	c := new(Config)

	if _, err := toml.DecodeFile(path, c); err != nil {
		return c, err
	}
	fmt.Println(c)

	return c, nil
}
