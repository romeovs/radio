package pulseaudio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseDevice(t *testing.T) {
	require := require.New(t)

	// Examples from playing around on my RPi
	tests := []struct {
		str string
		dev Device
	}{
		{
			str: "0	bt.monitor	module-null-sink.c	s16le 2ch 44100Hz	IDLE",
			dev: Device{
				Index:     0,
				Name:      "bt.monitor",
				Module:    "module-null-sink.c",
				Format:    "s16le",
				Channels:  2,
				Frequency: 44100,
				State:     StateIdle,
			},
		},
		{
			str: "1	alsa_output.platform-soc_audio.analog-mono.monitor	module-alsa-card.c	s16le 1ch 44100Hz	SUSPENDED",
			dev: Device{
				Index:     1,
				Name:      "alsa_output.platform-soc_audio.analog-mono.monitor",
				Module:    "module-alsa-card.c",
				Format:    "s16le",
				Channels:  1,
				Frequency: 44100,
				State:     StateSuspended,
			},
		},
		{
			str: "2	bluez_source.00_11_22_33_DE_D9.a2dp_source	module-bluez5-device.c	s16le 2ch 44100Hz RUNNING",
			dev: Device{
				Index:     2,
				Name:      "bluez_source.00_11_22_33_DE_D9.a2dp_source",
				Module:    "module-bluez5-device.c",
				Format:    "s16le",
				Channels:  2,
				Frequency: 44100,
				State:     StateRunning,
			},
		},
	}

	for _, tt := range tests {
		device, err := parseDevice(tt.str)
		require.Nil(err, "Parsing device string should return no errors")
		require.Equal(tt.dev.Index, device.Index, "Device index should be parsed correctly")
		require.Equal(tt.dev.Name, device.Name, "Device name should be parsed correctly")
		require.Equal(tt.dev.Module, device.Module, "Device module should be parsed correctly")
		require.Equal(tt.dev.Format, device.Format, "Device format should be parsed correctly")
		require.Equal(tt.dev.Channels, device.Channels, "Device channels should be parsed correctly")
		require.Equal(tt.dev.Frequency, device.Frequency, "Device frequency should be parsed correctly")
		require.Equal(tt.dev.State, device.State, "Device state should be parsed correctly")
	}
}
