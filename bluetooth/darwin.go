// +build darwin

package bluetooth

import (
	"fmt"
	"io"
)

func Bluetooth(name string) (io.ReadCloser, error) {
	return nil, fmt.Errorf("Bluetooth is not supported on this platform")
}
