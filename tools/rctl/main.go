package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/romeovs/radio/client"

	"github.com/voxelbrain/goptions"
)

type options struct {
	Help goptions.Help `goptions:"-h, --help, description='Show this help'"`
	Addr string        `goptions:"-a, --address, description='The address of the radio'"`
	goptions.Verbs

	Config    struct{} `goptions:"config"`
	SetConfig struct{} `goptions:"config:set"`

	Select struct{ goptions.Remainder } `goptions:"select"`
	Volume struct{ goptions.Remainder } `goptions:"volume"`
	Mute   struct{}                     `goptions:"mute"`
	UnMute struct{}                     `goptions:"unmute"`
}

func main() {
	opts := &options{
		Addr: "http://localhost:8080",
	}
	goptions.ParseAndFail(opts)

	client := client.New(opts.Addr)

	switch opts.Verbs {
	case "config":
		cfg, err := client.GetConfig()
		check(err)

		b, err := json.MarshalIndent(cfg, "", "  ")
		check(err)

		fmt.Println(string(b))

	case "config:set":
		notimplemented()

	case "select":
		args := opts.Select.Remainder
		if len(args) != 1 {
			check(errors.New("Expected 1 argument to select"))
		}

		ch, err := strconv.ParseUint(args[0], 10, 64)
		check(err)

		err = client.Select(int(ch))
		check(err)

	case "volume":
		args := opts.Volume.Remainder
		if len(args) != 1 {
			check(errors.New("Expected 1 argument to volume"))
		}

		ch, err := strconv.ParseUint(args[0], 10, 64)
		check(err)

		err = client.Volume(int(ch))
		check(err)

	case "mute":
		err := client.Mute()
		check(err)

	case "unmute":
		err := client.Unmute()
		check(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func notimplemented() {
	check(errors.New("not implemented"))
}
