package swap

import (
	"io"
	"sync"
)

// Reader is a TeeStream whose input can be changed on the fly.
type Reader struct {
	reader io.Reader
	lock   sync.RWMutex
}

// NewReader returns a new swappable reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{reader: r}
}

// Read reads data from the current stream in the swapper.
func (s *Reader) Read(p []byte) (int, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.reader == nil {
		return 0, nil
	}

	return s.reader.Read(p)
}

// Swap the current reader with r. The next call to Read will
// contain data from r instead of the previous reader.
func (s *Reader) Swap(r io.Reader) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.reader = r
}
