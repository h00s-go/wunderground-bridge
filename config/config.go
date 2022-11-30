package config

type Config struct {
	UpdateWunderground bool
}

func NewConfig(updateWunderground bool) (*Config, error) {
	return &Config{UpdateWunderground: updateWunderground}, nil
}
