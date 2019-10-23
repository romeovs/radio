package gpio

// PinState is the state of a pin
type PinState bool

const (
	// Low is the low state
	Low PinState = false

	// High is the high state
	High PinState = true
)

// State is the state of the whole GPIO setup.
type State []PinState

// Pin returns the state of the named pin.
func (s State) Pin(index int) PinState {
	return s[index-1]
}
