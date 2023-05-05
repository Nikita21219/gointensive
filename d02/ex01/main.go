package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type WC int

const (
	Lines WC = iota + 1
	Chars
	Words
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

func countWords(n int, data []byte) int {
	result := 0
	str := string(data[:n])
	lines := strings.Split(str, "\n")
	fmt.Println("lines: ", lines)
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			if word != "" {
				result++
			}
		}
	}
	return result
}

func readFile(path string, flag_ string) (string, error) {
	var result int

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	data := make([]byte, 5)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}

		if flag_ == "-w" {
			result += countWords(n, data)
		} else if flag_ == "-l" {
			result += strings.Count(string(data), "\n")
		} else if flag_ == "-m" {
			result += n
		}
	}

	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	if flag_ == "-w" {
	//		result += countWords(scanner.Text())
	//	} else if flag_ == "-l" {
	//		result += 1
	//	} else if flag_ == "-m" {
	//		result += len(scanner.Text()) + 1
	//	}
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	return "", err
	//}

	err = file.Close()
	return fmt.Sprintf("%d\t%s", result, path), err
}

func main() {

	//data := []byte("Hello Bold!")
	//file, err := os.Create("input3.txt")
	//if err != nil {
	//	fmt.Println("Unable to create file:", err)
	//	os.Exit(1)
	//}
	//defer file.Close()
	//for i := 0; i < 1000000; i++ {
	//	file.Write(data)
	//}
	//
	//fmt.Println("Done.")

	paths, flag_ := parseArgs()
	//fmt.Println("Flag:", flag_)
	//fmt.Println("Files:", paths)
	fmt.Println()

	for _, path := range paths {
		if result, err := readFile(path, flag_); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}
