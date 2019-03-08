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
	urlChannel := make(chan string, cfg.GetParallelRequestLimit())

	for _, rawURL := range cfg.GetURLs() {
		wg.Add(1)
		go func(url string) {
			urlChannel <- url
		}(rawURL)
	}

	go func() {
		wg.Wait()
		close(urlChannel)
	}()

	for rawURL := range urlChannel {
		agent := agent.NewAgent()
		response, err := agent.MakeRequest(rawURL)
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
		fmt.Printf("%v %v\n", rawURL, hashedBody)
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
