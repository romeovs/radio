package gpio

import (
	"github.com/stianeikeland/go-rpio"
)

// Selector polls the selector knob for changes.
//
// This expects the raspberry pi to be wired correctly with
// a Lorlin BCK1002 like so:
//
// BCK		RPI
// --------------
// GND		GND
// VCC		3.3v
// A			GPIO 22
// B			GPIO 23
// C			GPIO 24
// D			GPIO 25
//
type Selector struct {
	pins []rpio.Pin
}

var (
	selectorPinA = 6  // GPIO 22
	selectorPinB = 13 // GPIO 23
	selectorPinC = 19 // GPIO 24
	selectorPinD = 16 // GPIO 25
)

// NewSelector returns a new selector that starts listening right away.
func NewSelector() (*Selector, error) {
	pins := []rpio.Pin{
		rpio.Pin(selectorPinA),
		rpio.Pin(selectorPinB),
		rpio.Pin(selectorPinC),
		rpio.Pin(selectorPinD),
	}

	for _, pin := range pins {
		pin.Input()
		pin.PullDown()
		pin.Detect(rpio.AnyEdge)
	}

	return &Selector{pins: pins}, nil
}

// Read the pins and convert the output.
func (s *Selector) Read() int {
	return decodeSelector(
		s.pins[0].Read(),
		s.pins[1].Read(),
		s.pins[2].Read(),
		s.pins[3].Read(),
	)
}

// Close the selector and deregister all pins.
func (s *Selector) Close() {
	for _, pin := range s.pins {
		pin.Detect(rpio.NoEdge)
		pin.PullOff()
	}
}

// decodeSelector decodes the state of the rotary switch.
//
// See https://docs-emea.rs-online.com/webdocs/0ff6/0900766b80ff6d94.pdf
//
// In the version of the coded rotary switch I am using (Lorlin BCK1002), the code is given
// in hexadecimal binary compliment.
//
// This it the truth table from that document:
//
//          TERMINALS
//  VALUE   ABCD
//     1    1000
//     2    0100
//     3    1100
//     4    0010
//     5    1010
//     6    0110
//     7    1110
//     8    0001
//     9    1001
//    10    0101
//    11    1101
//    12    0011
//    13    1011
//    14    0111
//    15    1111
//    16    0000
func decodeSelector(A, B, C, D rpio.State) int {
	dail := 0b0000

	if A == rpio.High {
		dail = dail | 0b0001
	}

	if B == rpio.High {
		dail = dail | 0b0010
	}

	if C == rpio.High {
		dail = dail | 0b0100
	}

	if D == rpio.High {
		dail = dail | 0b1000
	}

	if dail == 0 {
		return 16
	}

	return dail
}
