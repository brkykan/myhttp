package coordinator

import (
	"agent"
	"config"
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// Run starts the execution of making requests with the given Configuration
func Run(cfg config.Configuration) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	coordinate(ctx, cfg)
}

func coordinate(ctx context.Context, cfg config.Configuration) {

	var wg sync.WaitGroup
	urlChannel := make(chan *url.URL, cfg.GetParallelRequestLimit())

	for _, rawURL := range cfg.GetURLs() {
		url, err := url.Parse(rawURL)
		if err != nil {
			log.Printf("Cannot parse URL %v: %v\n", url, err)
			continue
		}
		wg.Add(1)
		go func() {
			urlChannel <- url
		}()
	}

	go func() {
		wg.Wait()
		close(urlChannel)
	}()

	for url := range urlChannel {
		agent := agent.NewAgent()
		response, err := agent.MakeRequest(url)
		if err != nil {
			log.Printf("Error performing request: %+v\n", err)
			continue
		}
		body, err := getResponseBody(response)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			continue
		}
		hashedBody := hashResponse(body)
		fmt.Printf("%v %v\n", url.String(), hashedBody)
		wg.Done()
	}
}

func getResponseBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func hashResponse(body []byte) string {
	return fmt.Sprintf("%x", md5.Sum(body))
}
