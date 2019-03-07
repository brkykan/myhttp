package agent

import (
	"net/http"
	"net/url"
	"time"
)

type agent struct {
	client http.Client
}

type Requester interface {
	PerformGetRequest(url *url.URL) (*http.Response, error)
}

func (a *agent) PerformGetRequest(url *url.URL) (*http.Response, error) {
	if url.Scheme == "" {
		url.Scheme = "https"
	}
	resp, err := a.client.Get(url.String())
	return resp, err
}

func NewAgent() Requester {
	return &agent{
		client: http.Client{
			Timeout: 10 * time.Second, // a sensible time period to deem it as timed out
		},
	}
}
