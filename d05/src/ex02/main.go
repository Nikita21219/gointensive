package main

import (
	"container/heap"
	"fmt"
	"log"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap struct {
	presents []Present
}

func (h PresentHeap) Len() int { return len(h.presents) }

func (h PresentHeap) Less(i, j int) bool {
	if h.presents[i].Value == h.presents[j].Value {
		return h.presents[i].Size < h.presents[j].Size
	}
	return h.presents[i].Value > h.presents[j].Value
}

func (h PresentHeap) Swap(i, j int) { h.presents[i], h.presents[j] = h.presents[j], h.presents[i] }

func (h *PresentHeap) Push(x any) { (*h).presents = append((*h).presents, x.(Present)) }

func (h *PresentHeap) Pop() any {
	old := (*h).presents
	n := len((*h).presents)
	x := old[n-1]
	(*h).presents = old[0 : n-1]
	return x
}

func getNCoolestPresents(p []Present, n int) ([]Present, error) {
	var coolestPresents []Present
	if n > len(p) || n < 0 {
		return nil, fmt.Errorf("wrong argument 'n'")
	}

	h := &PresentHeap{}
	heap.Init(h)

	for _, pair := range p {
		heap.Push(h, pair)
	}

	for i := 0; i < n && len((*h).presents) > 0; i++ {
		coolestPresents = append(coolestPresents, heap.Pop(h).(Present))
	}
	return coolestPresents, nil
}

func main() {
	pair2 := Present{Value: 4, Size: 5}
	pair3 := Present{Value: 3, Size: 1}
	pair4 := Present{Value: 5, Size: 2}
	pair1 := Present{Value: 5, Size: 1}
	unsortedSliceOfPresents := []Present{pair1, pair2, pair3, pair4}

	presents, err := getNCoolestPresents(unsortedSliceOfPresents, 4)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, pair := range presents {
		fmt.Println(pair)
	}
}
