package bluetooth

import (
	"errors"
	"strings"

	"github.com/romeovs/radio/exec"
)

var (
	bluetoothDriver = "module-bluez5-device.c"

	// ErrNoBluetoothDevice is the error returned when no bluetooth device was found.
	ErrNoBluetoothDevice = errors.New("No Bluetooth device found")
)

// findSource tries to find the source of the bluetooth command.
func findSource() (string, error) {
	out, err := exec.Output("pactl", "list", "sources", "short")
	if err != nil {
		return "", err
	}

	return parseSources(string(out))
}

// parseSources parses the output from pactl and finds the first bluetooth source.
func parseSources(sources string) (string, error) {
	lines := strings.Split(sources, "\n")
	for _, line := range lines {
		flds := strings.Fields(line)
		if flds[2] == bluetoothDriver {
			return flds[0], nil
		}
	}

	return "", ErrNoBluetoothDevice
}
