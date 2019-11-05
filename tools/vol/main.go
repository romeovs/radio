// pot is a tool to test the potentiometer volume knob.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/romeovs/radio/gpio"
	rpio "github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	check(err)
	defer rpio.Close()

	vol, err := gpio.NewVolume()
	check(err)
	defer vol.Close()

	v := vol.Read()
	sig := interrupt()

	t := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-sig:
			t.Stop()
			return
		case <-t.C:
			vv := vol.Read()
			if vv != v {
				fmt.Println(vv)
				v = vv
			}
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func interrupt() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	return ch
}
