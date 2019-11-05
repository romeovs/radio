// sel is a tool to test the lorlin channel selector knob.
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

	sel, err := gpio.NewSelector()
	check(err)
	defer sel.Close()

	for {
		v := <-sel.Changes()
		fmt.Println(v)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
