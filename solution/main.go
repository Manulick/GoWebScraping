package main

import (
	"log"
	"sync"
)

func main() {
	log.Default().Println("Starting")
	response := Products{}
	wg := sync.WaitGroup{}

	wg.Add(len(urlList))
	for _, url := range urlList {
		go func(url string) {
			defer wg.Done()
			getUrlList(url, &response)
		}(url)

	}
	wg.Wait()

	createFile(response)
	log.Default().Println("Finish")
}
