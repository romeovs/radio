package web

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/romeovs/radio/convert"
)

// Stream the http web radio
func Stream(url string) (io.ReadCloser, error) {
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	typ := resp.Header.Get("Content-Type")
	if typ == "" {
		return nil, fmt.Errorf("Could not determine audio type from stream")
	}

	ext := strings.TrimPrefix(typ, "audio/")

	wav, err := convert.Convert(ext, resp.Body)
	if err != nil {
		return nil, err
	}

	return wav, nil
}
