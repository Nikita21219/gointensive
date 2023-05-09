package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func parseArgs() ([]string, string) {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("Wrong number of arguments")
	}

	var paths []string
	var flag_ string

	if len(args) >= 1 {
		flag_ = args[0]
		if flag_ != "-l" && flag_ != "-w" && flag_ != "-m" {
			flag_ = "-w"
			paths = args
		} else {
			paths = args[1:]
		}
	} else {
		flag_ = "-w"
	}

	return paths, flag_
}

func countWords(file *os.File) int {
	count := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}

func readFile(path string, flag_ string) (string, error) {
	var result int

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	if flag_ == "-w" {
		result += countWords(file)
	} else {
		data := make([]byte, 2048)

		for {
			n, err := file.Read(data)
			if err == io.EOF {
				break
			}
			if flag_ == "-l" {
				result += strings.Count(string(data), "\n")
			} else if flag_ == "-m" {
				result += n
			}
		}
	}

	return fmt.Sprintf("%d\t%s", result, path), file.Close()
}

func main() {
	//file, err := os.Create("/Users/a1/Downloads/IT/golang/gointensive/d02/ex01/input3.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//for i := 0; i < 100000000; i++ {
	//	data := []byte("Hello world\n")
	//	_, err = file.Write(data)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//}

	paths, flag_ := parseArgs()
	fmt.Println()
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)

		path := path
		go func() {
			defer wg.Done()
			result, err := readFile(path, flag_)
			if err != nil {
				fmt.Println("Error:", err)
				return
			} else {
				fmt.Println(result)
			}
		}()
	}
	wg.Wait()
}
