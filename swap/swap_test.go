package swap

import (
	"testing"
)

func TestSwap(t *testing.T) {
	as := &bytes{'a', 0}
	bs := &bytes{'b', 0}

	r := NewReader(as)

	str := ""

	// Read from the underlying reader and store the result in str
	read := func() {
		p := make([]byte, 2)

		n, err := r.Read(p)
		if err != nil {
			t.Errorf("Expected err = nil, got %s", err)
		}

		if n != chunk {
			t.Errorf("Expected n = 1, got %#v", chunk)
		}

		str += string(p[0:n])
	}

	read()
	if str != "a" {
		t.Errorf("Expected str = \"a\", got %#v", str)
	}

	read()
	if str != "aa" {
		t.Errorf("Expected str = \"aa\", got %#v", str)
	}

	r.Swap(bs)
	if as.closed != 1 {
		t.Errorf("Expected as.closed = 1, got %v", as.closed)
	}

	if bs.closed != 0 {
		t.Errorf("Expected bs.closed = 0, got %v", bs.closed)
	}

	read()
	if str != "aab" {
		t.Errorf("Expected str = \"aab\", got %#v", str)
	}

	read()
	if str != "aabb" {
		t.Errorf("Expected str = \"aabb\", got %#v", str)
	}

	r.Swap(as)
	if as.closed != 1 {
		t.Errorf("Expected as.closed = 1, got %v", as.closed)
	}

	if bs.closed != 1 {
		t.Errorf("Expected bs.closed = 1, got %v", bs.closed)
	}

	read()
	if str != "aabba" {
		t.Errorf("Expected str = \"aabba\", got %#v", str)
	}
}

func TestEmptySwap(t *testing.T) {
	r := NewReader(nil)
	p := make([]byte, 3)

	n, err := r.Read(p)
	if err != nil {
		t.Errorf("Expected err = nil, got %s", err)
	}

	if n != 0 {
		t.Errorf("Expected n = 0, got %v", n)
	}
}

// chunk the characters by this amount
var chunk = 1

// bytes implements a reader that reads the same byte always.
type bytes struct {
	b      byte
	closed int
}

// Read from bytes, returning the same bytes always.
func (b *bytes) Read(p []byte) (int, error) {
	n := len(p)
	if n > chunk {
		n = chunk
	}

	for i := 0; i < n; i++ {
		p[i] = b.b
	}

	return n, nil
}

// Close the stream, does nothing for testing purposes.
func (b *bytes) Close() error {
	b.closed += 1
	return nil
}
