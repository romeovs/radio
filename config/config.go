package config

// Config contains the user configuration of the radio.
type Config struct {
	// SayChannelName denotes wether or not the channel name should be
	// said by the radio when switching channels.
	SayChannelName bool `json:"say_channel_name"`

	// PlayStartupSound should denotes whether or not a startup sound should be played.
	PlayStartupSound bool `json:"startup_sound"`

	// Channels are all the channels.
	Channels []*Channel `json:"channels"`
}

// Channel is the config for a channel.
type Channel struct {
	// Type of the channel (webstream, ...)
	Type string `json:"type"`

	// Name of the channel, will be spoken out loud if SayChannelName is set to true.
	Name string `json:"name,omitempty"`

	// URL of the webstream or Spotify playlist.
	URL string `json:"url,omitempty"`

	// The name of the AirPlay device
	AirplayName string `json:"airplay_name,omitempty"`
}

// Channel gets the corresponding channel from the config.
func (cfg *Config) Channel(channel int) *Channel {
	var ch *Channel
	if channel >= 0 && channel < len(cfg.Channels) {
		ch = cfg.Channels[channel]
	}

	if ch == nil {
		return nil
	}

	return ch
}
