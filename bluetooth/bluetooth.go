package bluetooth

// See https://www.raspberrypi.org/forums/viewtopic.php?t=235519

import (
	"io"

	"github.com/romeovs/radio/exec"
)

// Bluetooth creates an audio stream that reads from bluetooth devices.
func Bluetooth(name string) (io.ReadCloser, error) {
	// Set the bluetooth alias
	err := setName(name)
	if err != nil {
		return nil, err
	}

	// Turn on bluetooth
	err = power(true)
	if err != nil {
		return nil, err
	}

	// Find the bluetooth device source name
	device, err := findSource()
	if err != nil {
		return nil, err
	}

	// Move the source output to bt
	// This assumes bt is a valid pulseaudio sink
	err = exec.Run("pactl", "move-source-output", "bt", device)
	if err != nil {
		return nil, err
	}

	// Listen to bt.monitor from pulseaudio
	r, err := exec.Command("parec", "-d", "bt.monitor")
	if err != nil {
		return nil, err
	}

	return &bluetoothWrapper{r}, nil
}

// bluetoothWrapper is a wrapper for io.ReadCloser that powers of the bluetooth module
// when closed.
type bluetoothWrapper struct {
	reader io.ReadCloser
}

// Read some bytes from the bluetooth stream.
func (r *bluetoothWrapper) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

// Close the bluetooth stream and power of the
func (r *bluetoothWrapper) Close() error {
	defer power(false)
	return r.reader.Close()
}

// Power bluetooth on or off.
func power(on bool) error {
	pwr := "off"
	if on {
		pwr = "on"
	}

	return exec.Run("bluetoothctl", "power", pwr)
}

// setName sets the bluetooth device name.
func setName(name string) error {
	return exec.Run("bluetoothctl", "system-alias", name)
}
