package exotel

import (
	"bytes"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "http://twilix.exotel.in/v1"
	// baseURL  = "http://requestb.in"
	rTimeout = 5 * time.Second
	cTimeout = 1 * time.Second
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

// Exotel : Holds the http client
type Exotel struct {
	Client *http.Client
	auth   Auth
}

// New : Get new Exotel type.
func New(userName string, password string) (e *Exotel, err error) {
	e = new(Exotel)
	err = e.auth.set(userName, password)
	if err != nil {
		return nil, err
	}
	e.setClient()
	return
}

// SetReadTimeout : Set read timeout for exotel request.
func (e *Exotel) SetReadTimeout(timeout time.Duration) {
	e.Client.Timeout = timeout
}

// SetConnectionTimeout : Sets connection timeout for TCP connection.
func (e *Exotel) SetConnectionTimeout(timeout time.Duration) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: timeout,
		}).Dial,
		TLSHandshakeTimeout: timeout,
	}
	e.Client.Transport = netTransport
}

func (e *Exotel) setClient() {
	e.Client = &http.Client{}
	e.SetReadTimeout(rTimeout)
	e.SetConnectionTimeout(cTimeout)
}

func (e *Exotel) doRequest(url string, params url.Values) (*http.Response, error) {
	exoRequest, err := http.NewRequest("POST", url, bytes.NewBufferString(params.Encode()))
	if err != nil {
		return nil, err
	}
	exoRequest.SetBasicAuth(e.auth.Username, e.auth.Password)
	exoRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := e.Client.Do(exoRequest)
	if err != nil {
		return nil, err
	}
	return response, nil
}
