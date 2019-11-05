package gpio

import (
	"testing"

	"github.com/stianeikeland/go-rpio"
)

func TestDecodeSelector(t *testing.T) {
	// See the truth table in the comments on decodeSelector
	testDecodeSelector(t, 0b1000, 1)
	testDecodeSelector(t, 0b0100, 2)
	testDecodeSelector(t, 0b1100, 3)
	testDecodeSelector(t, 0b0010, 4)
	testDecodeSelector(t, 0b1010, 5)
	testDecodeSelector(t, 0b0110, 6)
	testDecodeSelector(t, 0b1110, 7)
	testDecodeSelector(t, 0b0001, 8)
	testDecodeSelector(t, 0b1001, 9)
	testDecodeSelector(t, 0b0101, 10)
	testDecodeSelector(t, 0b1101, 11)
	testDecodeSelector(t, 0b0011, 12)
	testDecodeSelector(t, 0b1011, 13)
	testDecodeSelector(t, 0b0111, 14)
	testDecodeSelector(t, 0b1111, 15)
	testDecodeSelector(t, 0b0000, 16)
}

// toSelectorState is a helper that turns a binary representation of the pin state into
// a usable State.
func toSelectorState(bin int) [4]rpio.State {
	s := [4]rpio.State{}

	if bin&0b1000 != 0 {
		s[0] = rpio.High
	}

	if bin&0b0100 != 0 {
		s[1] = rpio.High
	}

	if bin&0b0010 != 0 {
		s[2] = rpio.High
	}

	if bin&0b0001 != 0 {
		s[3] = rpio.High
	}

	return s
}

// testDecodeSelector tests whether or not decodeSelector decoding works properly.
func testDecodeSelector(t *testing.T, bin, expected int) {
	s := toSelectorState(bin)
	v := decodeSelector(s[0], s[1], s[2], s[3])
	if expected != v {
		t.Errorf("Expected state %#v to result in value %v, but got %v", s, expected, v)
	}
}
