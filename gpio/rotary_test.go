package gpio

import (
	"testing"

	"github.com/stianeikeland/go-rpio"
)

func TestRotatrySwitch(t *testing.T) {
	// See the truth table in the comments on rotarySwitch
	testRotarySwitch(t, 0b1000, 1)
	testRotarySwitch(t, 0b0100, 2)
	testRotarySwitch(t, 0b1100, 3)
	testRotarySwitch(t, 0b0010, 4)
	testRotarySwitch(t, 0b1010, 5)
	testRotarySwitch(t, 0b0110, 6)
	testRotarySwitch(t, 0b1110, 7)
	testRotarySwitch(t, 0b0001, 8)
	testRotarySwitch(t, 0b1001, 9)
	testRotarySwitch(t, 0b0101, 10)
	testRotarySwitch(t, 0b1101, 11)
	testRotarySwitch(t, 0b0011, 12)
	testRotarySwitch(t, 0b1011, 13)
	testRotarySwitch(t, 0b0111, 14)
	testRotarySwitch(t, 0b1111, 15)
	testRotarySwitch(t, 0b0000, 16)
}

// toSwitchState is a helper that turns a binary representation of the pin state into
// a usable State.
func toSwitchState(bin int) [4]rpio.State {
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

// TestRotarySwitch tests wether or not rotaerySwitch decoding works properly.
func testRotarySwitch(t *testing.T, bin, expected int) {
	s := toSwitchState(bin)
	v := rotarySwitch(s[0], s[1], s[2], s[3])
	if expected != v {
		t.Errorf("Expected state %#v to result in value %v, but got %v", s, expected, v)
	}
}
