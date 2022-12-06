package application

type Station struct {
	Weather *Weather
}

func NewStation() *Station {
	return &Station{
		Weather: &Weather{},
	}
}
