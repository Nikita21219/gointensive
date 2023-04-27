package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type DBReader interface {
	Read()
}

func parseArgs() (string, bool) {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Wrong number of arguments")
		return "", true
	}
	split := strings.Split(args[0], ".")
	if len(split) != 2 {
		fmt.Println("Wrong argument")
		return "", true
	}
	if split[1] == "xml" || split[1] == "json" {
		return args[0], false
	} else {
		fmt.Println("Wrong argument")
		return "", true
	}
}

type ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []ingredient `json:"ingredients" xml:"ingredients"`
}

type recipes struct {
	Cake []cake `json:"cake" xml:"cake"`
}

func readJson(filePath string) {
	b := getDataFromFile(filePath)
	var r recipes

	if !json.Valid(b) {
		fmt.Println("Not valid JSON")
		os.Exit(1)
	}
	err := json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println("Error!")
	}
	printXml(r)
}

func getDataFromFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("File does not exists")
		os.Exit(1)
	}
	return data
}

func printJson(r recipes) {
	data, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		fmt.Println("Fatal error")
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func printXml(r recipes) {
	data, err := xml.MarshalIndent(r, "", "    ")
	if err != nil {
		fmt.Println("Fatal error")
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func main() {
	filePath, err := parseArgs()
	if err {
		return
	}
	switch strings.Split(filePath, ".")[1] {
	case "xml":
		fmt.Println("Test")
	case "json":
		readJson(filePath)
	}
}
