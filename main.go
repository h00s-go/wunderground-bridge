package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/h00s-go/wunderground-bridge/application"
	"github.com/h00s-go/wunderground-bridge/config"
	"github.com/h00s-go/wunderground-bridge/mqtt"
)

func main() {
	cfg, err := config.NewConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var m *mqtt.MQTT
	if cfg.MQTT.Enabled {
		m = mqtt.NewMQTT(&cfg.MQTT)
		defer m.Close()
	}

	app := application.NewApplication(cfg, logger, m)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", app.NewRouter())
}
