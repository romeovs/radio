package pulseaudio

import "fmt"

// State is the state of a pulseaudio device.
type State string

var (
	// StateRunning is the running state.
	StateRunning = State("RUNNING")

	// StateIdle is the idle state.
	StateIdle = State("IDLE")

	// StateSuspended is the suspended state.
	StateSuspended = State("SUSPENDED")
)

func parseState(str string) (State, error) {
	switch str {
	case "RUNNING":
		return StateRunning, nil
	case "IDLE":
		return StateIdle, nil
	case "SUSPENDED":
		return StateSuspended, nil
	default:
		return State("UNKNOWN"), fmt.Errorf("Unknown state %s", str)
	}
}

// String implements fmt.Stringer for State.
func (s State) String() string {
	return string(s)
}
