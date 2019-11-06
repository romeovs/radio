package audio

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

const defaultSink = "@DEFAULT_SINK@"

// Play the wav audio from the stream
func Play(s io.Reader) error {
	play := exec.Command("paplay", "--raw", "--latency-msec=300")
	play.Stdin = s
	play.Stderr = os.Stderr

	return play.Run()
}

// Volume sets the system volume in percentage (0-100).
func Volume(volume uint) error {
	if volume > 100 {
		volume = 100
	}

	cmd := exec.Command("pactl", "set-sink-volume", defaultSink, fmt.Sprintf("%v%%", volume))
	return cmd.Run()
}

// Mute or unmute the volume.
func Mute(mute bool) error {
	m := "0"
	if mute {
		m = "1"
	}

	cmd := exec.Command("pactl", "set-sink-mute", defaultSink, m)
	return cmd.Run()
}
