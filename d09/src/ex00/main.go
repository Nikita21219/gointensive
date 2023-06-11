package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func notValidArray(nums []int) bool {
	for _, n := range nums {
		if n < 0 {
			return true
		}
	}
	return false
}

func sleepSort(nums []int) (chan int, error) {
	if notValidArray(nums) {
		return nil, fmt.Errorf("Not valid array")
	}
	wg := sync.WaitGroup{}
	ch := make(chan int, len(nums))

	for _, n := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}(n)
	}

	wg.Wait()
	close(ch)

	return ch, nil
}

func main() {
	ch, err := sleepSort([]int{3, 2, 1, 4})
	if err != nil {
		log.Fatalln(err)
	}
	for num := range ch {
		fmt.Println(num)
	}
}
