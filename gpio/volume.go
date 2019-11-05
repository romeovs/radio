package gpio

import (
	"math"
	"sync"
	"time"

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
	curr int
	done bool
	ch   chan int
	lock sync.RWMutex
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

	v := &Volume{
		curr: read(),
		ch:   make(chan int),
		done: false,
	}

	go func() {
		for {
			if ok := v.poll(); !ok {
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return v, nil
}

// poll is one step in the polling process.
// Returns true if it should keep on polling, false if the polling should end.
func (v *Volume) poll() bool {
	v.lock.Lock()
	defer v.lock.Unlock()

	if v.done {
		return false
	}

	vol := read()
	if vol != v.curr {
		v.curr = vol
		v.ch <- vol
	}

	return true
}

// Changes returns a channel on which changes in volume will be post.
func (v *Volume) Changes() <-chan int {
	return v.ch
}

// Close the Volume reader.
func (v *Volume) Close() {
	v.lock.Lock()
	defer v.lock.Unlock()

	v.done = true
	close(v.ch)

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