// pot is a tool to test the potentiometer volume knob.
//
// Wire up an MCP3008 to the raspberry pi like so:
//
// MCP		RPI
// -----------
// VDD		3.3V
// VREF		3.3V
// AGND		GND
// CLK		GPIO11
// DOUT		GPIO9
// DIN		GPIO10
// CS			GPIO8
// DGND		GND
//
package main

import (
	"fmt"
	"log"
	"math"
	"time"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

var (
	channel = 0
)

func main() {
	err := rpio.Open()
	check(err)
	defer rpio.Close()

	err = rpio.SpiBegin(rpio.Spi0)
	check(err)
	defer rpio.SpiEnd(rpio.Spi0)

	rpio.SpiSpeed(1350000)
	rpio.SpiChipSelect(0)

	for {
		buf := []byte{1, byte((8 + channel) << 4), 0}
		rpio.SpiExchange(buf)
		fmt.Printf("%v%%\n", toVolume(buf))
		time.Sleep(200 * time.Millisecond)
	}
}

func toVolume(buf []byte) int {
	// Create a value 0 - 1023
	V := int(buf[1]&3)<<8 + int(buf[2])
	vol := (1023 - float64(V)) / 10.23

	return int(math.Round(vol))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
