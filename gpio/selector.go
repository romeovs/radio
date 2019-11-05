package gpio

import (
	"sync"
	"time"

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
	done bool
	curr int
	ch   chan int
	lock sync.Mutex
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
	s := &Selector{
		ch: make(chan int),
		pins: []rpio.Pin{
			rpio.Pin(selectorPinA),
			rpio.Pin(selectorPinB),
			rpio.Pin(selectorPinC),
			rpio.Pin(selectorPinD),
		},
	}

	for _, pin := range s.pins {
		pin.Input()
		pin.PullDown()
		pin.Detect(rpio.AnyEdge)
	}

	s.curr = s.read()

	go func() {
		for {
			if ok := s.poll(); !ok {
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return s, nil

}

// read the pins and convert the output.
func (s *Selector) read() int {
	return rotarySwitch(
		s.pins[0].Read(),
		s.pins[1].Read(),
		s.pins[2].Read(),
		s.pins[3].Read(),
	)
}

// poll for changes, returns false if we should stop polling,
// and true if otherwise.
func (s *Selector) poll() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.done {
		return false
	}

	ch := s.read()
	if ch != s.curr {
		s.curr = ch
		s.ch <- ch
	}

	return true
}

// Close the selector and deregister all pins.
func (s *Selector) Close() {
	for _, pin := range s.pins {
		pin.Detect(rpio.NoEdge)
		pin.PullOff()
	}

	s.done = true
}

// Changes returns a channel that contains the changes in the pin.
func (s *Selector) Changes() <-chan int {
	return s.ch
}
