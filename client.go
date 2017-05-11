package exotel

import (
	"io"
	"net"
	"net/http"
	"time"

	"github.com/devpyp/exotel/exoerror"
)

const (
	baseURL  = "http://twilix.exotel.in/v1"
	rTimeout = 5 * time.Second
	cTimeout = 1 * time.Second
)

const (
	timeFormat = "2006-01-02 15:04:05 MST"
)

const (
	post = "POST"
	get  = "GET"
)

// Exotel : Holds the http client
type Exotel struct {
	Client *http.Client
	auth   Auth
}

// New : Get new Exotel type.
func New(userName string, password string) (e *Exotel, err *exoerror.Error) {
	e = new(Exotel)
	er := e.auth.set(userName, password)
	if er != nil {
		err = exoerror.AuthError
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

func (e *Exotel) doRequest(method string, url string, body io.Reader) (res *http.Response, err *exoerror.Error) {
	exoRequest, er := http.NewRequest(method, url, body)
	if er != nil {
		err = exoerror.ClientError
		return
	}
	exoRequest.SetBasicAuth(e.auth.Username, e.auth.Password)
	exoRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, er = e.Client.Do(exoRequest)
	if er != nil || res.StatusCode != 200 {
		err = exoerror.RequestFailed
		return
	}
	return
}
