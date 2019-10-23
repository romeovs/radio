package main

import (
	"flag"
	"log"

	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/http"
	"github.com/romeovs/radio/radio"
)

var (
	// configFile is the location of the config file.
	configFile = flag.String("config", "./radio.json", "the location of the config file")
)

func main() {
	flag.Parse()

	cfg, err := config.Open(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Build radio.
	radio := radio.NewRadio(cfg)

	// Set up http server.
	go http.New(radio).Listen(":8080")

	// Start the radio.
	radio.Start()
}
