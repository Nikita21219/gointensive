package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func handleSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	for {
		switch <-sigCh {
		case os.Interrupt:
			cancel()
			return
		}
	}
}

func getBody(url string) *string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error GET request:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		return nil
	}
	bodyStr := string(body)
	return &bodyStr
}

func crawlWeb(ctx context.Context, ch <-chan string) chan *string {
	resCh := make(chan *string)

	go func() {
		defer close(resCh)

		sem := make(chan struct{}, 8)

		wg := sync.WaitGroup{}
		for url := range ch {
			url := url
			wg.Add(1)
			sem <- struct{}{}
			go func() {
				defer wg.Done()
				defer func() { <-sem }()

				select {
				case <-ctx.Done():
					return
				case resCh <- getBody(url):

				}
			}()
		}
		wg.Wait()
	}()

	return resCh
}

func main() {
	urlsList := []string{
		"https://youtube.com",
		"https://ya.ru",
		"https://habr.com/",
		"https://madorsky.site/",
		"https://github.com/Nikita21219",
		"https://edu.21-school.ru/",
		"https://translate.yandex.ru/",
		"https://translate.yandex.ru/",
		"https://drive.google.com/",
		"https://google.com/",
		"https://www.deepl.com/translator",
		"https://www.speedtest.net/",
	}
	urls := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go handleSignals(cancel)

	// Fill url channel
	go func() {
		for _, u := range urlsList {
			urls <- u
		}
		close(urls)
	}()

	res := crawlWeb(ctx, urls)

	for r := range res {
		fmt.Println("Result: ", *r)
	}
}
