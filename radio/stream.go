package radio

import (
	"io"

	"github.com/romeovs/radio/airplay"
	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/web"
)

// stream creates the channels stream.
func stream(channel *config.Channel) (io.Reader, error) {
	switch channel.Type {
	case "webradio":
		return web.Stream(channel.URL)
	case "airplay":
		return airplay.Airplay(channel.AirplayName)
	}
	return nil, nil
}
