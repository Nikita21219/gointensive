package main

import (
	"fmt"
	"sync"
)

func asChan(vs ...interface{}) chan interface{} {
	ch := make(chan interface{})

	go func() {
		for _, v := range vs {
			ch <- v
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}

func multiplex(channels ...<-chan interface{}) <-chan interface{} {
	outCh := make(chan interface{})
	wg := sync.WaitGroup{}

	for _, c := range channels {
		c := c
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				outCh <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

func main() {
	a := asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := asChan(10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	c := asChan(20, 21, 22, 23, 24, 25, 26, 27, 28, 29)
	outChannel := multiplex(a, b, c)

	for v := range outChannel {
		fmt.Println(v)
	}
}
