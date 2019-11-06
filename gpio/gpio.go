package gpio

import (
	"github.com/romeovs/radio/radio"
)

// GPIO monitors changes on the IO pins of a Raspberry Pi.
type GPIO struct {
	radio *radio.Radio
	vol   *Volume
	sel   *Selector
	amp   *AMP2
	done  chan bool
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

// Mute or unmute the AMP2.
func (g *GPIO) Mute(mute bool) {
	if g.amp != nil {
		g.amp.Mute(mute)
	}
}

// change the channel.
func (g *GPIO) change(channel int) {
	_ = g.radio.Select(channel)
}

// volume sets the volume in the underlying radio.
func (g *GPIO) volume(vol int) {
	_ = g.radio.Volume(uint(vol))
}
