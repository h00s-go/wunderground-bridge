package station

import "github.com/h00s-go/wunderground-bridge/config"

type Wunderground struct {
	Config *config.Wunderground
}

func NewWunderground(config *config.Wunderground) *Wunderground {
	return &Wunderground{
		Config: config,
	}
}
