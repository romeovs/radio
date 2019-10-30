// +build linux

package gpio

import (
	"time"

	"github.com/romeovs/radio/log"
	rpio "github.com/stianeikeland/go-rpio/v4"
)

var (
	// delay for polling the pins.
	delay = 300 * time.Millisecond

	// the total number of pins.
	total = 22

	// an array containing the pins.
	pins = make([]rpio.Pin, total)
)

func init() {
	// Set up all the pins
	for i := range pins {
		pins[i] = rpio.Pin(i + 1)
	}
}

func (g *GPIO) setup() {
	err := rpio.Open()
	if err != nil {
		log.Error("ERROR SETTING UP rpio: %s", err)
	}
	defer rpio.Close()

	// Mark all pins as input, pull them down and start detecting edges.
	for _, pin := range pins {
		pin.Input()
		pin.PullDown()
		pin.Detect(rpio.AnyEdge)
	}

	state := make([]PinState, len(pins))
	for i := range pins {
		state[i] = false
	}

	for {
		time.Sleep(delay)

		changed := false

		for i, pin := range pins {
			if pin.EdgeDetected() {
				changed = true
				state[i] = PinState(pin.Read() == rpio.High)
			}
		}

		if changed {
			g.event(State(state))
		}
	}
}

func (g *GPIO) event(state State) {
	// TODO: implement this
	log.Info("EVENT %#v", state)
}
