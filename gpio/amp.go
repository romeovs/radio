package gpio

import "github.com/stianeikeland/go-rpio"

var mutePin = 23 // GPIO 4

// AMP2 interfaces with the Hifiberry AMP2 hat.
type AMP2 struct {
	pin rpio.Pin
}

// NewAMP2 creates a new AMP2.
func NewAMP2() (*AMP2, error) {
	return &AMP2{
		pin: rpio.Pin(mutePin),
	}, nil
}

// Mute the AMP2 by disabling the output stage.
func (a *AMP2) Mute(mute bool) {
	if mute {
		a.pin.PullDown()
	} else {
		a.pin.PullOff()
	}
}

// Close the AMP2 and reset io pins.
func (a *AMP2) Close() {
	a.pin.PullOff()
}
