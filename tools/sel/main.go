// sel is a tool to test the lorlin channel selector knob.
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

	sel, err := gpio.NewSelector()
	check(err)
	defer sel.Close()

	s := sel.Read()

	sig := interrupt()
	t := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-t.C:
			ss := sel.Read()
			if ss != s {
				fmt.Println(ss)
				s = ss
			}
		case <-sig:
			t.Stop()
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
