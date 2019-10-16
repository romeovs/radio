package audio

import (
	"io"
	"os"
	"os/exec"
)

// Play the wav audio from the stream
func Play(s io.Reader) error {
	play := exec.Command(cmd)
	play.Stdin = s
	play.Stderr = os.Stderr

	return play.Run()
}
