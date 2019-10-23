package sounds

import (
	"bytes"
	"io"
)

// Startup returns the startup sound.
func Startup() io.Reader {
	return bytes.NewReader(_bindataSoundsWavStartupwav)
}
