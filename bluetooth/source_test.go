package bluetooth

import (
	"strings"
	"testing"
)

func TestParseSources(t *testing.T) {
	// Example from playing around on my RPi
	sources := strings.Join([]string{
		"0	bt.monitor	module-null-sink.c	s16le 2ch 44100Hz	IDLE",
		"1	alsa_output.platform-soc_audio.analog-mono.monitor	module-alsa-card.c	s16le 1ch 44100Hz	SUSPENDED",
		"2	bluez_source.00_11_22_33_DE_D9.a2dp_source	module-bluez5-device.c	s16le 2ch 44100HzSUSPENDED",
	}, "\n")

	source, err := parseSources(sources)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if source != "2" {
		t.Errorf("Expected bluetooth device to be \"2\" but got, \"%s\"", source)
	}
}

func TestParseSourcesNoDevice(t *testing.T) {
	// Example from playing around on my RPi
	sources := strings.Join([]string{
		"0	bt.monitor	module-null-sink.c	s16le 2ch 44100Hz	IDLE",
		"1	alsa_output.platform-soc_audio.analog-mono.monitor	module-alsa-card.c	s16le 1ch 44100Hz	SUSPENDED",
	}, "\n")

	source, err := parseSources(sources)
	if err != ErrNoBluetoothDevice {
		t.Errorf("Expected ErrNoBluetoothDevice, got %v", err)
	}

	if source != "" {
		t.Errorf("Expected bluetooth device to be \"\" but got, \"%s\"", source)
	}
}
