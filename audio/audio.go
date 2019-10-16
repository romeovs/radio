package audio

import (
	"io"
	"os"
	"os/exec"
)

// Play the wav audio from the stream
func Play(s io.Reader) error {
	play := exec.Command("paplay")
	play.Stdin = s
	play.Stderr = os.Stderr

	return play.Run()
}
