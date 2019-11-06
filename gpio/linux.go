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

	g.amp, err = NewAMP2()
	if err != nil {
		log.Error("ERROR SETTING UP AMP2: %s", err)
	}

	g.vol, err = NewVolume()
	if err != nil {
		log.Error("ERROR SETTING UP VOLUME KNOB: %s", err)
	}

	g.sel, err = NewSelector()
	if err != nil {
		log.Error("ERROR SETTING UP CHANNEL KNOB: %s", err)
	}

	vol := cache{g.vol.Read()}
	sel := cache{g.sel.Read()}

	t := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case _, ok := <-g.done:
			t.Stop()
			return
		case <-t.C:
			if v, changed := vol.Changed(g.vol.Read()); changed {
				g.volume(v)
			}

			if s, changed := sel.Changed(g.sel.Read()); changed {
				g.change(s)
			}
		}
	}
}

// Close closes the GPIO object.
func (g *GPIO) Close() {
	rpio.Close()
	g.vol.Close()
	g.sel.Close()
	g.amp.Close()

	close(g.done)
}

type cache struct {
	value int
}

func (c *cache) Changed(newvalue int) (int, bool) {
	res := c.value == newvalue
	c.value = newvalue
	return
}
