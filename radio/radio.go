package radio

import (
	"fmt"
	"io"

	"github.com/romeovs/radio/audio"
	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/sounds"
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
func (r *Radio) Select(channel int) error {
	ch := r.Config.Channel(channel)

	if ch == nil {
		// TODO: log error
		fmt.Printf("NO CHANNEL [%v] FOUND\n", channel)
		return fmt.Errorf("Channel %v does not exist", channel)
	}

	fmt.Printf("SELECTING [%v] \"%s\"\n", channel, ch.Name)
	if r.Config.SayChannelName {
		t, err := speech.Say(ch.Name)
		if err != nil {
			fmt.Println("ERROR SPEAKING CHANNEL NAME \"%s\": %s", ch.Name, err)
			return fmt.Errorf("Error saying channel name stream: %s", err)
		}

		r.playall(t)
	}

	s, err := stream(ch)
	if err != nil {
		fmt.Println("ERROR OPENING STREAM:", err)
		return fmt.Errorf("Error opening stream: %s", err)
	}

	r.play(s)
	return nil
}

// play swaps the underlying stream to the new stream
func (r *Radio) play(s io.Reader) {
	r.stream.Swap(s, false)
}

// playall swaps the underlying stream to the new stream
// and makes sure it gets played fully.
func (r *Radio) playall(s io.Reader) {
	r.stream.Swap(s, true)
}

// Start starts playing the audio stream.
func (r *Radio) Start() {
	if r.Config.PlayStartupSound {
		r.playall(sounds.Startup())
	}

	audio.Play(r.stream)
}

// Volume sets the radio volume.
func (r *Radio) Volume(volume uint) error {
	fmt.Printf("SETTING VOLUME %v%%\n", volume)
	return audio.Volume(volume)
}
