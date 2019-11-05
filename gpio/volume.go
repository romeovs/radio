package gpio

import (
	"math"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

// Volume can read the status of the volume knob.
//
// This expects the raspberry pi to be wired correctly with
// an MCP3008 like so:
//
// MCP		RPI
// -----------
// VDD		3.3V
// VREF		3.3V
// AGND		GND CLK		GPIO11
// DOUT		GPIO9
// DIN		GPIO10
// CS			GPIO8
// DGND		GND
//
type Volume struct {
}

// NewVolume creates a new volume that is ready for use.
// This expects rpio.Open to be called already.
func NewVolume() (*Volume, error) {
	err := rpio.SpiBegin(rpio.Spi0)
	if err != nil {
		return nil, err
	}

	rpio.SpiSpeed(1350000)
	rpio.SpiChipSelect(0)

	return &Volume{}, nil
}

// Read the value of the volume knob.
func (v *Volume) Read() int {
	return read()
}

// Close the Volume reader.
func (v *Volume) Close() {
	rpio.SpiEnd(rpio.Spi0)
}

// read reads the volume from the SPI module.
func read() int {
	buf := []byte{1, byte(8 << 4), 0}
	rpio.SpiExchange(buf)
	return MCP3008Volume(buf)
}

// MCP3008Volume converts an MCP3008 voltage reading into
// a volume 0% - 100%.
func MCP3008Volume(buf []byte) int {
	// Create a value 0 - 1023
	h := int(buf[1]&0b11) << 8
	l := int(buf[2])
	V := h + l
	vol := 100 * (1023 - float64(V)) / 1023

	return int(math.Round(vol))
}
