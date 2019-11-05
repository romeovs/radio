package gpio

import (
	"github.com/romeovs/radio/radio"
)

// GPIO monitors changes on the IO pins of a Raspberry Pi.
type GPIO struct {
	radio *radio.Radio
	vol   *Volume
	sel   *Selector
}

// New returns a new GPIO instance.
func New(radio *radio.Radio) *GPIO {
	return &GPIO{
		radio: radio,
	}
}

// Listen starts listen to changes on the GPIO pins.
func (g *GPIO) Listen() {
	g.setup()
}

// change the channel.
func (g *GPIO) change(channel int) {
	g.radio.Select(channel)
}

// volume sets the volume in the underlying radio.
func (g *GPIO) volume(vol int) {
	g.radio.Volume(uint(vol))
}
