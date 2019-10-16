package speech

import (
	"io"

	"github.com/romeovs/radio/exec"
)

// Say the text.
func Say(text string) (io.ReadCloser, error) {
	return exec.Command("espeak", text, "--stdout")
}
