package main

import (
	"io"
	"log"

	"github.com/romeovs/radio/audio"
	"github.com/romeovs/radio/speech"
	"github.com/romeovs/radio/web"
)

func main() {
	text, err := speech.Say("NTS Radio")
	check(err)

	stream, err := web.Stream("https://stream-relay-geo.ntslive.net/stream")
	check(err)

	multi := io.MultiReader(text, stream)

	err = audio.Play(multi)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
