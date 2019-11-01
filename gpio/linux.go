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
)

func (g *GPIO) setup() {
	err := rpio.Open()
	if err != nil {
		log.Error("ERROR SETTING UP rpio: %s", err)
	}
	defer rpio.Close()

	vol, err := NewVolume()
	if err != nil {
		log.Error("ERROR SETTING UP VOLUME KNOB: %s", err)
	}
	defer vol.Close()

	vch := vol.Changes()

	for {
		select {
		case v := <-vch:
			g.volume(v)
		}
	}
}
