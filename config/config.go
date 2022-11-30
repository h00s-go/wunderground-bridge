package config

type Config struct {
	UpdateWunderground    bool
	WundergroundUpdateURL string
}

func NewConfig(updateWunderground bool) (*Config, error) {
	return &Config{
		UpdateWunderground:    updateWunderground,
		WundergroundUpdateURL: "http://rtupdate.wunderground.com/weatherstation/updateweatherstation.php",
	}, nil
}
