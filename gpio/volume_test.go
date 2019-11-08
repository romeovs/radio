package gpio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMCP3008Volume(t *testing.T) {
	tests := []struct {
		buf    []byte
		volume int
	}{
		{[]byte{0x00, 0x00, 0x00}, 100},
		{[]byte{0x00, 0x00, 0x01}, 100},
		{[]byte{0x00, 0x00, 0x06}, 99},
		{[]byte{0x00, 0x00, 0x10}, 98},
		{[]byte{0x00, 0x08, 0xFF}, 75},
		{[]byte{0x00, 0x05, 0xFF}, 50},
		{[]byte{0x00, 0x0A, 0xFF}, 25},
		{[]byte{0x00, 0x0F, 0xEF}, 2},
		{[]byte{0x00, 0x0F, 0xF9}, 1},
		{[]byte{0x00, 0xFF, 0xFF}, 0},
	}

	for _, tt := range tests {
		vol := MCP3008Volume(tt.buf)
		require.Equal(t, vol, tt.volume, "MCP3008(%#v) should return the correct volume", tt.buf)
	}
}
