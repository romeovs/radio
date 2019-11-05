// sel is a tool to test the lorlin channel selector knob.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

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

	sel, err := gpio.NewSelector()
	check(err)
	defer sel.Close()

	sig := interrupt()

	for {
		select {
		case v := <-sel.Changes():
			fmt.Println(v)
		case <-sig:
			return
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
