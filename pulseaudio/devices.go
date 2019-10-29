package pulseaudio

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/romeovs/radio/exec"
)

// Device is a pulseaudio device (source or sink).
type Device struct {
	// Index is the device index.
	Index int

	// Name is the name given to the device.
	Name string

	// Module is the pulseaudio module that is responsible for making the device work.
	Module string

	// Format is the audio format of the device.
	Format string

	// Channels is the amount of channels the device has.
	Channels int

	// Frequency is the audio frequency of the device.
	Frequency int

	// State is the state the device is currently in.
	State State
}

// ListSources lists all currently available sources.
func ListSources() ([]Device, error) {
	out, err := exec.Output("pactl", "list", "sources", "short")
	if err != nil {
		return nil, err
	}

	return parseDevices(string(out))
}

// ListSinks lists all currently available sinks.
func ListSinks() ([]Device, error) {
	out, err := exec.Output("pactl", "list", "sinks", "short")
	if err != nil {
		return nil, err
	}

	return parseDevices(string(out))
}

// parseDevices parses the output from a pactl list _ short command and
// returns the devices (sources/sinks) it found.
func parseDevices(out string) ([]Device, error) {
	lines := strings.Split(out, "\n")
	devices := make([]Device, len(lines))

	for i, line := range lines {
		dev, err := parseDevice(line)
		if err != nil {
			return nil, err
		}

		devices[i] = dev
	}

	return devices, nil
}

// parseDevice parses a single line oft the output from a pactl list _ short command and
// and returns the device (source/sink) info it found.
func parseDevice(line string) (Device, error) {
	flds := strings.Fields(line)

	index, err := strconv.ParseUint(flds[0], 10, 64)
	if err != nil {
		return Device{}, fmt.Errorf("Expected integer for device index but got \"%s\"", flds[0])
	}

	channels, err := strconv.ParseUint(flds[4][0:len(flds[4])-2], 10, 64)
	if err != nil {
		return Device{}, fmt.Errorf("Expected integer for device channels but got \"%s\"", flds[4])
	}

	frequency, err := strconv.ParseUint(flds[5][0:len(flds[5])-2], 10, 64)
	if err != nil {
		return Device{}, fmt.Errorf("Expected integer for device frequency but got \"%s\"", flds[4])
	}

	state, err := parseState(flds[6])
	if err != nil {
		return Device{}, err
	}

	return Device{
		Index:     int(index),
		Name:      flds[1],
		Module:    flds[2],
		Format:    flds[3],
		Channels:  int(channels),
		Frequency: int(frequency),
		State:     state,
	}, nil
}
