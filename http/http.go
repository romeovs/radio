package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kataras/muxie"
	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/radio"
)

// Server is a http server that serves the HTTP API for the radio.
type Server struct {
	mux   *muxie.Mux
	radio *radio.Radio
}

// New creates a new Server.
func New(radio *radio.Radio) *Server {
	mux := muxie.NewMux()

	server := &Server{
		mux:   mux,
		radio: radio,
	}

	server.setup()

	return server
}

// setup sets up the http muxer.
func (s *Server) setup() {
	s.mux.Handle("/config",
		muxie.Methods().
			HandleFunc("GET", s.handleGetConfig).
			HandleFunc("PUT", s.handleSetConfig),
	)

	s.mux.Handle("/select/:channel",
		muxie.Methods().
			HandleFunc("PUT", s.handleSelect),
	)

	s.mux.Handle("/volume/:percentage",
		muxie.Methods().
			HandleFunc("PUT", s.handleSetVolume),
	)
}

// ServeHTTP implements http.Server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// handleGetConfig returns the config of the radio.
func (s *Server) handleGetConfig(w http.ResponseWriter, r *http.Request) {
	err := muxie.Dispatch(w, muxie.JSON, s.radio.Config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot marshal json: %s", err), http.StatusInternalServerError)
	}
}

// handleSetConfig updates the config of the radio.
func (s *Server) handleSetConfig(w http.ResponseWriter, r *http.Request) {
	cfg := new(config.Config)
	err := muxie.Bind(r, muxie.JSON, cfg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot parse json: %s", err), http.StatusBadRequest)
	}

	s.radio.Config = cfg

	w.WriteHeader(http.StatusNoContent)
}

// handleSelect selects a radio channel.
func (s *Server) handleSelect(w http.ResponseWriter, r *http.Request) {
	param := muxie.GetParam(w, "channel")
	channel, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Expected integer for channel but got '%s'", param), http.StatusBadRequest)
	}

	err = s.radio.Select(int(channel))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) handleSetVolume(w http.ResponseWriter, r *http.Request) {
	param := muxie.GetParam(w, "percentage")
	volume, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Expected integer for volume but got '%s'", param), http.StatusBadRequest)
	}

	err = s.radio.Volume(uint(volume))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Listen to the address and serve the server there.
func (s *Server) Listen(addr string) {
	http.ListenAndServe(addr, s)
}
