package main

import (
	"fmt"
	"log"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if idx < 0 || idx >= len(arr) {
		return -1, fmt.Errorf("Wrong index")
	}

	ptr := unsafe.Pointer(&arr[0])
	res := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(idx)*unsafe.Sizeof(ptr)))
	return *res, nil
}

func main() {
	el, err := getElement([]int{1, 4, 34, 43, 435, 3}, 2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(el)
}
