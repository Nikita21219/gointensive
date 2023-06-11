package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func parseArgs() ([]string, map[string]bool) {
	charFlag := flag.Bool("m", false, "chars")
	wordsFlag := flag.Bool("w", false, "words")
	linesFlag := flag.Bool("l", false, "lines")
	flag.Parse()

	paths := make([]string, 0, 5)

	i := 0
	args := os.Args[1:]
	for idx, arg := range args {
		if strings.HasPrefix(arg, "-") {
			i = idx + 1
			continue
		}
		break
	}
	for k := i; k < len(args); k++ {
		paths = append(paths, args[k])
	}

	if len(paths) == 0 {
		log.Fatalln("Not found files")
	}

	return paths, map[string]bool{
		"chars": *charFlag,
		"words": *wordsFlag,
		"lines": *linesFlag,
	}
}

func countWords(data []byte) int {
	count := 0
	reader := bufio.NewReader(bytes.NewReader(data))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	return count
}

func readFile(path string, flags map[string]bool) (int, int, int, error) {
	var resultWords int
	var resultChars int
	var resultLines int

	file, err := os.Open(path)
	if err != nil {
		return 0, 0, 0, err
	}

	data := make([]byte, 2048+4)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, 0, 0, err
		}

		if flags["lines"] {
			resultLines += strings.Count(string(data[:n]), "\n")
		}
		if flags["chars"] {
			resultChars += utf8.RuneCount(data[:n])
		}
		if flags["words"] {
			resultWords += countWords(data[:n])
		}
	}

	_ = file.Close()

	return resultWords, resultChars, resultLines, nil
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

	paths, flags := parseArgs()
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)

		path := path
		go func() {
			defer wg.Done()
			resultWords, resultChars, resultLines, err := readFile(path, flags)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("\t")
			if flags["lines"] {
				fmt.Printf("%d\t", resultLines)
			}
			if flags["words"] {
				fmt.Printf("%d\t", resultWords)
			}
			if flags["chars"] {
				fmt.Printf("%d\t", resultChars)
			}
			fmt.Println("\t", path)
		}()
	}
	wg.Wait()
}
