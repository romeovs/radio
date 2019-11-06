package http

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/kataras/muxie"
	"github.com/romeovs/radio/config"
	"github.com/romeovs/radio/log"
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

	s.mux.Handle("/volume/:volume",
		muxie.Methods().
			HandleFunc("PUT", s.handleSetVolume),
	)

	s.mux.Handle("/mute/:mute",
		muxie.Methods().
			HandleFunc("PUT", s.handleSetMute),
	)

	s.mux.Handle("/logs",
		muxie.Methods().
			HandleFunc("GET", s.handleGetLogs),
	)
}

// ServeHTTP implements http.Server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug("HTTP API: %s %s", r.Method, r.URL.String())
	s.mux.ServeHTTP(w, r)
}

// handleGetConfig returns the config of the radio.
func (s *Server) handleGetConfig(w http.ResponseWriter, r *http.Request) {
	err := muxie.Dispatch(w, muxie.JSON, s.radio.Config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot marshal json: %s", err), http.StatusInternalServerError)
		return
	}
}

// handleSetConfig updates the config of the radio.
func (s *Server) handleSetConfig(w http.ResponseWriter, r *http.Request) {
	cfg := new(config.Config)
	err := muxie.Bind(r, muxie.JSON, cfg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot parse json: %s", err), http.StatusBadRequest)
		return
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
		return
	}

	err = s.radio.Select(int(channel))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleSetVolume(w http.ResponseWriter, r *http.Request) {
	param := muxie.GetParam(w, "volume")
	volume, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Expected integer for volume but got '%s'", param), http.StatusBadRequest)
		return
	}

	err = s.radio.Volume(uint(volume))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleSetMute(w http.ResponseWriter, r *http.Request) {
	param := muxie.GetParam(w, "mute")

	if param != "true" && param != "false" {
		http.Error(w, fmt.Sprintf("Expected 0 or 1 for mute but got '%s'", param), http.StatusBadRequest)
		return
	}

	err := s.radio.Mute(param == "true")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleGetLogs(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile(log.LogFile, os.O_RDONLY, 0666)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer f.Close()

	if r.URL.Query().Get("plain") == "1" {
		buf := bufio.NewReaderSize(f, 64*1024)

		for {
			entry := new(log.Entry)
			line, isPrefix, err := buf.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				http.Error(w, fmt.Sprintf("Cannot read log file: %s", err), http.StatusInternalServerError)
				return
			}

			if isPrefix {
				// TODO: write line to w
				w.Write(line)
				w.Write([]byte("\n"))
				continue
			}

			err = json.NewDecoder(bytes.NewBuffer(line)).Decode(entry)
			if err != nil {
				http.Error(w, fmt.Sprintf("Cannot parse log file: %s", err), http.StatusInternalServerError)
				return
			}

			log.Fprint(w, entry)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	io.Copy(w, f)
}

// Listen to the address and serve the server there.
func (s *Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s)
}
