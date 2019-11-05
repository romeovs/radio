package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/gpio"
	"github.com/romeovs/radio/http"
	"github.com/romeovs/radio/log"
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
		log.Error("Error opening config file: %s", err)
		os.Exit(1)
	}

	defer func() {
		fmt.Println()
		log.Info("BYE")
	}()

	// Build radio.
	radio := radio.NewRadio(cfg)

	// Set up http server.
	go http.New(radio).Listen(":8080")

	// Set up gpio pins.
	io := gpio.New(radio)
	go io.Listen()
	defer io.Close()

	// Start the radio.
	go radio.Start()

	// Wait for interrupt.
	wait()
}

func wait() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
