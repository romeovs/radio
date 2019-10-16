package config

import (
	"io"
)

// Config contains the user configuration of the radio.
type Config struct {
	chans []Channel
}

// A Channel is a channel that can be played by the radio player.
type Channel interface {
	Stream() io.Reader
}

// Channel gets the corresponding channel from the config.
func (cfg *Config) Channel(channel int) Channel {
	var ch Channel
	if channel >= 0 && channel < len(cfg.chans) {
		ch = cfg.chans[channel]
	}

	if ch == nil {
		return &Silence{}
	}

	return ch
}
