package airplay

import (
	"io"

	"github.com/romeovs/radio/exec"
)

// Airplay starts an airplay server and pipes the audio.
func Airplay(name string) (io.ReadCloser, error) {
	if name == "" {
		name = "Radio"
	}

	return exec.Command("shairport-sync", "-a", name, "-o", "stdout")
}
