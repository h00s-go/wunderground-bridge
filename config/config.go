package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Wunderground Wunderground
	MQTT         MQTT
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

func NewConfig(path string) (*Config, error) {
	c := new(Config)

	if _, err := toml.DecodeFile(path, c); err != nil {
		return c, err
	}

	return c, nil
}
