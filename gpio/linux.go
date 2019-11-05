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

	g.vol, err = NewVolume()
	if err != nil {
		log.Error("ERROR SETTING UP VOLUME KNOB: %s", err)
	}

	for {
		select {
		case v := <-vol.Changes():
			g.volume(v)
		}
	}
}

// Close closes the GPIO object.
func (g *GPIO) Close() {
	rpio.Close()
	g.volume.Close()
}
