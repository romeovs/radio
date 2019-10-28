package swap

import (
	"io"
	"sync"
	"time"
)

// Reader is a TeeStream whose input can be changed on the fly.
type Reader struct {
	// reader is the current reader that is playing.
	reader io.Reader

	// lock is used for syncing
	lock sync.RWMutex

	// wg can be used to wait on the end of the current reader.
	wg *sync.WaitGroup

	// done is true if the current stream has returned io.EOF at least once.
	done bool

	// all is true if the current stream must be played until the end as
	// specified by the Swap all parameter.
	all bool
}

// NewReader returns a new swappable reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		reader: r,
	}
}

// Read reads data from the current stream in the swapper.
func (s *Reader) Read(p []byte) (int, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.reader == nil {
		return 0, nil
	}

	n, err := s.reader.Read(p)

	if err == io.EOF {
		if !s.done {
			s.wg.Done()
			s.done = true
		}

		time.Sleep(2 * time.Millisecond)
		return 0, nil
	}

	return n, err
}

// Swap the current reader with r.
// Future calls to Read will contain data from r instead of the previous reader.
//
// If the current reader can be closed (ie implements io.Closer), it will be.
//
// Set the all parameter to to true if you want to make sure r will be used until io.EOF
// before other readers are allowed to be swapped in.
func (s *Reader) Swap(r io.Reader, all bool) {
	// Wait for previous stream to end
	s.wait()

	// Lock for writing
	s.lock.Lock()
	defer s.lock.Unlock()

	// Close the current reader
	if closer, ok := s.reader.(io.Closer); ok {
		closer.Close()
	}

	s.reader = r
	s.all = all

	s.wg = new(sync.WaitGroup)
	s.wg.Add(1)
	s.done = false
}

func (s *Reader) wait() {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.all && s.wg != nil {
		s.wg.Wait()
	}
}
