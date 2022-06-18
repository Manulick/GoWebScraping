package main

import (
	"GoWebScraping/solution/openapi"
	"log"
	"sync"
)

func main() {
	log.Default().Println("Starting")
	response := openapi.Response{}
	wg := sync.WaitGroup{}

	wg.Add(len(urlList))
	for _, url := range urlList {
		go func(url string) {
			defer wg.Done()
			getUrlList(url, &response)
		}(url)

	}
	wg.Wait()
	log.Default().Println("Finish")
}
