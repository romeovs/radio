package convert

import (
	"io"

	"github.com/romeovs/radio/exec"
)

var args = []string{
	"-acodec", "pcm_s16le",
	"-f", "s16le",
	"-ac", "2",
	"-ar", "44000",
	"pipe:.wav",
}

// Convert a stream of the specified type in to wav.
func Convert(typ string, stream io.ReadCloser) (io.ReadCloser, error) {
	input := "pipe:." + typ
	a := []string{"-i", input}
	return exec.Pipe(stream, "ffmpeg", append(a, args...)...)
}
