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

	var wg sync.WaitGroup
	wg.Add(len(cfg.GetURLs()))
	go pool(ctx, &wg, len(cfg.GetURLs()), cfg)
	wg.Wait()
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

func worker(tasksCh <-chan string, wg *sync.WaitGroup, workerNumber int) {
	defer wg.Done()

	for {
		rawURL, ok := <-tasksCh
		if !ok {
			return
		}

		agent := agent.NewAgent()
		response, err := agent.MakeRequest(rawURL)
		wg.Done()
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
		fmt.Printf("Worker number %d produced %v %v\n", workerNumber, rawURL, hashedBody)
		wg.Done()
	}
}

func pool(ctx context.Context, wg *sync.WaitGroup, workers int, cfg config.Configuration) {
	tasksCh := make(chan string)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg, i)
	}

	for i := 0; i < len(cfg.GetURLs()); i++ {
		tasksCh <- cfg.GetURLs()[i]
	}

	close(tasksCh)
}
