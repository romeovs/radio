package config

import (
	"io"

	"github.com/romeovs/radio/audio"
)

// Silence is a Channel that is silent.
type Silence struct{}

// Stream returns the silent stream.
func (s *Silence) Stream() io.Reader {
	return audio.Silence
}

// Read reads zeroes into p.
func (s *silentReader) Read(p []byte) (int, error) {
	n := cap(p)

	// Chunk the reads
	if n > 4608 {
		n = 4608
	}

	for i = 0; i < n; i++ {
		p[i] = 0
	}

	return n, nil
}
