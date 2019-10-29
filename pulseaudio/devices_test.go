package pulseaudio

import (
	"testing"
)

func TestParseDevice(t *testing.T) {
	{
		// Example from playing around on my RPi
		str := "0	bt.monitor	module-null-sink.c	s16le 2ch 44100Hz	IDLE"

		device, err := parseDevice(str)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		deviceEquals(t, device, Device{
			Index:     0,
			Name:      "bt.monitor",
			Module:    "module-null-sink.c",
			Format:    "s16le",
			Channels:  2,
			Frequency: 44100,
			State:     StateIdle,
		})
	}

	{
		// Example from playing around on my RPi
		str := "1	alsa_output.platform-soc_audio.analog-mono.monitor	module-alsa-card.c	s16le 1ch 44100Hz	SUSPENDED"

		device, err := parseDevice(str)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		deviceEquals(t, device, Device{
			Index:     1,
			Name:      "alsa_output.platform-soc_audio.analog-mono.monitor",
			Module:    "module-alsa-card.c",
			Format:    "s16le",
			Channels:  1,
			Frequency: 44100,
			State:     StateSuspended,
		})
	}

	{
		// Example from playing around on my RPi
		str := "2	bluez_source.00_11_22_33_DE_D9.a2dp_source	module-bluez5-device.c	s16le 2ch 44100Hz RUNNING"

		device, err := parseDevice(str)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}

		deviceEquals(t, device, Device{
			Index:     2,
			Name:      "bluez_source.00_11_22_33_DE_D9.a2dp_source",
			Module:    "module-bluez5-device.c",
			Format:    "s16le",
			Channels:  2,
			Frequency: 44100,
			State:     StateRunning,
		})
	}
}

func deviceEquals(t *testing.T, device, expected Device) {
	if expected.Index != device.Index {
		t.Errorf("Expected device.Index to be %v, got %v", expected.Index, device.Index)
	}

	if expected.Name != device.Name {
		t.Errorf("Expected device.Name to be \"%s\" got \"%s\"", expected.Name, device.Name)
	}

	if expected.Module != device.Module {
		t.Errorf("Expected device.Module to be \"%s\" got \"%s\"", expected.Module, device.Module)
	}

	if expected.Format != device.Format {
		t.Errorf("Expected device.Format to be \"%s\" got \"%s\"", expected.Format, device.Format)
	}

	if expected.Channels != device.Channels {
		t.Errorf("Expected device.Channels to be %v got %v", expected.Channels, device.Channels)
	}

	if expected.Frequency != device.Frequency {
		t.Errorf("Expected device.Frequency to be %v got %v", expected.Frequency, device.Frequency)
	}

	if expected.State != device.State {
		t.Errorf("Expected device.State to be %s got %s", expected.State, device.State)
	}
}
