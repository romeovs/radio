package gpio

import (
	"github.com/romeovs/radio/radio"
)

// GPIO monitors changes on the IO pins of a Raspberry Pi.
type GPIO struct {
	radio *radio.Radio
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

func (g *GPIO) change(channel int) {
	g.radio.Select(channel)
}
