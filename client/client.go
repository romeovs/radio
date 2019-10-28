package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/romeovs/radio/config"
)

// Client talks to a running radio using the HTTP API.
type Client struct {
	addr string
}

// New returns a new client for the radio at the specified address.
func New(addr string) Client {
	return Client{addr}
}

func request(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	if check(resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func check(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return fmt.Errorf(string(body))
}

func (c Client) url(format string, args ...interface{}) string {
	return c.addr + fmt.Sprintf(format, args...)
}

// GetConfig gets the radios configuration.
func (c Client) GetConfig() (*config.Config, error) {
	resp, err := request(http.MethodGet, c.url("/config"))
	if err != nil {
		return nil, err
	}

	v := &config.Config{}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// Select the specified channel on the radio.
func (c Client) Select(channel int) error {
	_, err := request(http.MethodPut, c.url("/select/%v", channel))
	return err
}

// Volume sets the specified volume on the radio.
func (c Client) Volume(volume int) error {
	_, err := request(http.MethodPut, c.url("/volume/%v", volume))
	return err
}

// Mute the radio.
func (c Client) Mute() error {
	_, err := request(http.MethodPut, c.url("/mute/true"))
	return err
}

// Unmute the radio.
func (c Client) Unmute() error {
	_, err := request(http.MethodPut, c.url("/mute/false"))
	return err
}
