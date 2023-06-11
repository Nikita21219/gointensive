package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getSortedDataFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error open file: %s", err)
	}
	defer file.Close()
	var data []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	sort.Slice(data, func(i, j int) bool {
		return strings.Compare(data[i], data[j]) < 0
	})

	return data
}

func remove(idx int, arr *[]string) {
	*arr = append((*arr)[:idx], (*arr)[idx+1:]...)
}

func printDiff(str string, arr *[]string) {
	idx := binarySearch(*arr, str)
	if idx == len(*arr) {
		if str != "" {
			fmt.Printf("ADDED %s\n", str)
		}
	} else {
		remove(idx, arr)
	}
}

func readFileAndCompare(data []string, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		printDiff(scanner.Text(), &data)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	for _, str := range data {
		fmt.Printf("REMOVED %s\n", str)
	}
}

func binarySearch(arr []string, key string) int {
	left := 0
	right := len(arr)
	var mid int

	for left <= right {
		mid = left + (right-left)/2
		if mid >= len(arr) {
			return len(arr)
		}
		if key < arr[mid] {
			right = mid - 1
		} else if key > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return len(arr)
}

func main() {
	var oldFile string
	var newFile string
	flag.StringVar(&oldFile, "old", "", "Need to use '--old' flag")
	flag.StringVar(&newFile, "new", "", "Need to use '--new' flag")
	flag.Parse()

	if oldFile == "" || newFile == "" {
		fmt.Println("Wrong arguments")
		return
	}

	data := getSortedDataFromFile(oldFile)
	readFileAndCompare(data, newFile)
}
