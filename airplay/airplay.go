package airplay

import (
	"io"
	"time"

	"github.com/romeovs/radio/exec"
)

// Airplay starts an airplay server and pipes the audio.
func Airplay(name string) (io.ReadCloser, error) {
	if name == "" {
		name = "Radio"
	}

	r, err := exec.Command("shairport-sync", "-a", name, "-o", "stdout")
	if err != nil {
		return nil, err
	}

	return NewRaceReader(r), nil
}

// RaceReader is a reader that will return an empty read
// after a while if nothing is returned from the underlying reader.
type RaceReader struct {
	reader io.ReadCloser

	reads chan []byte
}

// NewRaceReader wraps the reader in a reader that will always return
// at least something from the Read function.
func NewRaceReader(r io.ReadCloser) *RaceReader {
	race := &RaceReader{
		reader: r,
		reads:  make(chan []byte, 1),
	}

	go func() {
		// TODO: How to end this go function?
		for {
			p := make([]byte, 512)
			n, _ := r.Read(p)
			race.reads <- p[:n]
		}
	}()

	return race
}

func (r *RaceReader) Read(p []byte) (int, error) {
	select {
	case buf := <-r.reads:
		for i, b := range buf {
			p[i] = b
		}

		return len(buf), nil
	case <-time.After(400 * time.Millisecond):
		return 0, nil
	}
}

// Close closes the reader.
func (r *RaceReader) Close() error {
	// close(r.reads)
	return r.reader.Close()
}
