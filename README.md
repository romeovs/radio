# radio

Radio is a simple program that turns a Raspberry Pi into a radio that can play
audio from a number of different sources.


## Sources

These are the sources that have been implemented for now:

- Web stream
- Airplay
- Bluetooth

## Run locally 

Run Radio from your shell:

```sh
go cmd/main.go
```

Set a channel using curl, e.g.
 
```sh
curl -XPUT http://localhost:8080/select/1
 ```

See `http.go` for more

## Dependencies

Radio uses PulseAudio. If you are on a non-linux OS you will need to install it, for example via Homebrew:

```sh
brew install pulseaudio
```

And start it:

```sh
brew services start pulseaudio
```

## Configuration

The radio can be configured via a config file (default `radio.json` or can be
set using the `-c` flag). See [`radio.json`](radio.json) for an example configuration.


## Hardware

In the future I will post hardware I've used to build the radio here, as well as
schema's on how to wire up the GPIO for the control knobs.
