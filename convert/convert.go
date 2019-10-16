package convert

import (
	"io"

	"github.com/romeovs/radio/exec"
)

// Convert a stream of the specified type in to wav.
func Convert(typ string, stream io.ReadCloser) (io.ReadCloser, error) {
	input := "pipe:." + typ
	return exec.Pipe(stream, "ffmpeg", "-i", input, "pipe:.wav")
}
