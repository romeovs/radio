package gpio

import "testing"

func TestMCP3008Volume(t *testing.T) {
	type test struct {
		buf    []byte
		volume int
	}

	tests := []test{
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

	for _, test := range tests {
		r := MCP3008Volume(test.buf)
		if r != test.volume {
			t.Errorf("Expected MCP3008Volume(%#v) to be %v, but got %v", test.buf, test.volume, r)
		}
	}
}
