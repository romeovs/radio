package bluetooth

import (
	"errors"
	"fmt"

	"github.com/romeovs/radio/pulseaudio"
)

var (
	bluetoothDriver = "module-bluez5-device.c"

	// ErrNoBluetoothDevice is the error returned when no bluetooth device was found.
	ErrNoBluetoothDevice = errors.New("No Bluetooth device found")
)

// findSource tries to find the source of the bluetooth command.
func findSource() (string, error) {
	sources, err := pulseaudio.ListSources()
	if err != nil {
		return "", err
	}

	for _, source := range sources {
		if source.Module == bluetoothDriver {
			return fmt.Sprintf("%v", source.Index), nil
		}
	}

	return "", ErrNoBluetoothDevice
}
