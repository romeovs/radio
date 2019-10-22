package radio

import (
	"fmt"
	"io"

	"github.com/romeovs/radio/audio"
	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/speech"
	"github.com/romeovs/radio/swap"
)

// Radio is the main radio model
type Radio struct {
	Config *config.Config

	stream *swap.Reader
}

// NewRadio creates a new radio.
func NewRadio(cfg *config.Config) *Radio {
	return &Radio{
		Config: cfg,
		stream: swap.NewReader(nil),
	}
}

// Select a radio channel
// This closes the previously playing channel
func (r *Radio) Select(channel int) {
	fmt.Printf("Trying %v\n", channel)
	ch := r.Config.Channel(channel)

	if ch == nil {
		// TODO: log error
		fmt.Println("No channel found")
		return
	}

	if r.Config.SayChannelName {
		if t, _ := speech.Say(ch.Name); t != nil {
			// TODO: handle error
			r.play(t)
		}
	}

	fmt.Printf("Selecting %v: %s\n", channel, ch.Name)
	s, err := stream(ch)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	r.play(s)
}

// play swaps the underlying stream to the new stream
func (r *Radio) play(s io.Reader) {
	r.stream.Swap(s)
}

// Start starts playing the audio stream.
func (r *Radio) Start() {
	audio.Play(r.stream)
}
