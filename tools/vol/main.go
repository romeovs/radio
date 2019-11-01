// pot is a tool to test the potentiometer volume knob.
package main

import (
	"fmt"
	"log"

	"github.com/romeovs/radio/gpio"
	rpio "github.com/stianeikeland/go-rpio/v4"
)

var (
	channel = 0
)

func main() {
	err := rpio.Open()
	check(err)
	defer rpio.Close()

	vol, err := gpio.NewVolume()
	check(err)
	defer vol.Close()

	changes := vol.Changes()

	for {
		v := <-changes
		fmt.Println(v)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
