package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type WC int

const (
	Lines WC = iota + 1
	Chars
	Words
)

func parseArgs() ([]string, WC) {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("Wrong number of arguments")
	}
	paths := make([]string, 0)

	linesFlag := flag.Bool("l", false, "for counting lines")
	charFlag := flag.Bool("m", false, "for counting characters")
	wordsFlag := flag.Bool("w", false, "for counting words")
	flag.Parse()

	for _, arg := range os.Args {
		fmt.Println("ARG: ", arg)
	}

	counter := 0
	for _, ptr := range []*bool{linesFlag, charFlag, wordsFlag} {
		if *ptr {
			counter++
		}
	}
	if counter == 0 {
		*wordsFlag = true
	} else if counter > 1 {
		log.Fatalln("Wrong flags")
	}

	var flag_ WC
	if *linesFlag {
		flag_ = Lines
	} else if *charFlag {
		flag_ = Chars
	} else if *wordsFlag {
		flag_ = Words
	}
	return paths, flag_
}

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error to open file:", path)
		return
	}
	fmt.Println(file)
}

func main() {
	path, flag_ := parseArgs()
	fmt.Println(path)
	fmt.Println(flag_)
	//readFile("Hello")
}
