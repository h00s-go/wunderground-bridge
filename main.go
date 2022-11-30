package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/h00s-go/wunderground-bridge/application"
	"github.com/h00s-go/wunderground-bridge/config"
)

func main() {
	cfg, err := config.NewConfig(true)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := application.NewApplication(cfg, logger)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", app.NewRouter())
}
