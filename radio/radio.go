package radio

import (
	"fmt"
	"io"

	"github.com/romeovs/radio/audio"
	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/log"
	"github.com/romeovs/radio/sounds"
	"github.com/romeovs/radio/speech"
	"github.com/romeovs/radio/swap"
)

// Radio is the main radio model
type Radio struct {
	Config *config.Config
	OnMute func(mute bool)

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
		log.Info("NO CHANNEL [%v] FOUND", channel)
		return fmt.Errorf("Channel %v does not exist", channel)
	}

	log.Info("SELECTING [%v] \"%s\"", channel, ch.Name)
	if r.Config.SayChannelName {
		t, err := speech.Say(ch.Name)
		if err != nil {
			log.Error("ERROR SPEAKING CHANNEL NAME \"%s\": %s", ch.Name, err)
			return fmt.Errorf("Error saying channel name stream: %s", err)
		}

		r.playall(t)
	}

	s, err := stream(ch)
	if err != nil {
		log.Error("ERROR OPENING STREAM: %s", err)
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
	log.Info("WELCOME")
	if r.Config.PlayStartupSound {
		r.playall(sounds.Startup())
	}

	err := audio.Play(r.stream)
	if err != nil {
		log.Error("CANNOT PLAY STREAM: %s", err)
	}
}

// Volume sets the radio volume.
func (r *Radio) Volume(volume uint) error {
	if volume > 100 {
		return fmt.Errorf("INVALID VOLUME VALUE %v", volume)
	}

	log.Info("SETTING VOLUME %v%%", volume)
	return audio.Volume(volume)
}

// Mute mutes or unmutes the radio.
func (r *Radio) Mute(mute bool) error {
	log.Info("MUTING %v", mute)

	if r.OnMute != nil {
		r.OnMute(mute)
	}

	return audio.Mute(mute)
}
