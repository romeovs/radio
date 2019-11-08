package swap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// chunk the characters by this amount
var chunk = 1

func TestSwap(t *testing.T) {
	require := require.New(t)

	as := &bytes{'a', 0}
	bs := &bytes{'b', 0}

	r := NewReader(as)

	str := ""

	// Read from the underlying reader and store the result in str
	read := func() {
		p := make([]byte, 2)

		n, err := r.Read(p)
		require.Nil(err, "reading from a swapper with a reader should not return nil")
		require.Equal(n, chunk, "reading from a swapper should return all available bytes")
		str += string(p[0:n])
	}

	read()
	require.Equal(str, "a", "string should be \"a\"")

	read()
	require.Equal(str, "aa", "string should be \"aa\"")

	r.Swap(bs, false)
	require.Equal(as.closed, 1, "as should be closed once after first Swap")
	require.Equal(bs.closed, 0, "bs should not be closed after first Swap")

	read()
	require.Equal(str, "aab", "string should be \"aab\"")

	read()
	require.Equal(str, "aabb", "string should be \"aabb\"")

	r.Swap(as, false)
	require.Equal(as.closed, 1, "as should be closed once after second Swap")
	require.Equal(bs.closed, 1, "bs should be closed once after second Swap")

	read()
	require.Equal(str, "aabba", "string should be \"aabba\"")
}

func TestEmptySwap(t *testing.T) {
	require := require.New(t)

	r := NewReader(nil)
	p := make([]byte, 3)

	n, err := r.Read(p)
	require.Nil(err, "reading from an empty swapper should not return an error")
	require.Zero(n, "reading from an empty swapper return 0 bytes")
}

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
