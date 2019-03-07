package coordinator

import (
	"agent"
	"config"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
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

	fmt.Println(cfg.GetURLs())

	for _, rawURL := range cfg.GetURLs() {
		url, err := url.Parse(rawURL)
		if err != nil {
			log.Printf("Cannot parse URL %v: %v\n", url, err)
			continue
		}
		wg.Add(1)
		fmt.Printf("sending url %v into channel\n", url.String())
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
		fmt.Printf("requesting %v\n", url)
		response, err := agent.PerformGetRequest(url)
		if err != nil {
			log.Printf("Error performing request: %+v\n", err)
			return
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return
		}
		hash := md5.New()
		hashed := hash.Sum(body)
		md5String := hex.EncodeToString(hashed)
		fmt.Printf("%v %v\n", url, md5String)
		wg.Done()
	}
}

func printHash(md5String, url string) {
}
