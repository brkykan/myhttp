package agent

import (
	"net/http"
	"strings"
	"time"
)

type Agent struct {
	client http.Client
}

// Requester in an interface that implements MakeRequest function
type Requester interface {
	MakeRequest(url string) (*http.Response, error)
}

func (a *Agent) MakeRequest(url string) (*http.Response, error) {

	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		url = strings.TrimPrefix(url, "https://")
		url = strings.TrimPrefix(url, "http://")
		url = "http://" + url
	}
	resp, err := a.client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// NewAgent creates and returns a Requester
func NewAgent() Requester {
	return &Agent{
		client: http.Client{
			Timeout: 10 * time.Second, // a sensible time period to deem it as timed out
		},
	}
}
