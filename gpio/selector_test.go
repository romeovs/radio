package gpio

import (
	"testing"

	"github.com/stianeikeland/go-rpio"
)

func TestDecodeSelector(t *testing.T) {
	// See the truth table in the comments on decodeSelector
	tests := []struct {
		pins int
		out  int
	}{
		{0b1000, 1},
		{0b0100, 2},
		{0b1100, 3},
		{0b0010, 4},
		{0b1010, 5},
		{0b0110, 6},
		{0b1110, 7},
		{0b0001, 8},
		{0b1001, 9},
		{0b0101, 10},
		{0b1101, 11},
		{0b0011, 12},
		{0b1011, 13},
		{0b0111, 14},
		{0b1111, 15},
		{0b0000, 16},
	}

	for _, tt := range tests {
		v := decodeSelector(
			toState(tt.pins, 0b1000),
			toState(tt.pins, 0b0100),
			toState(tt.pins, 0b0010),
			toState(tt.pins, 0b0001),
		)

		if tt.out != v {
			t.Errorf("Expected state %04b to result in value %v, but got %v", tt.pins, tt.out, v)
		}
	}
}

func toState(s int, mask int) rpio.State {
	if s&mask != 0 {
		return rpio.High
	}
	return rpio.Low
}
