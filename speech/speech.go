package speech

import (
	"io"

	"github.com/romeovs/radio/convert"
	"github.com/romeovs/radio/exec"
)

// Say the text.
func Say(text string) (io.ReadCloser, error) {
	r, err := exec.Command("espeak", text, "--stdout", "-p", "40", "-s", "160")
	if err != nil {
		return nil, err
	}

	return convert.Convert("wav", r)
}
