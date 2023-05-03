package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type DBReader interface {
	Read(filePath string) (Recipes, error)
}

type JsonReader struct{}

type XmlReader struct{}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Recipes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

func (j *JsonReader) Read(filePath string) (Recipes, error) {
	b := getDataFromFile(filePath)
	var r Recipes

	if !json.Valid(b) {
		err := fmt.Errorf("Not valid JSON")
		return r, err
	}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (j *XmlReader) Read(filePath string) (Recipes, error) {
	b := getDataFromFile(filePath)
	var r Recipes

	err := xml.Unmarshal(b, &r)
	if err != nil {
		return r, err
	}
	return r, nil
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

func getDataFromFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("File does not exists")
		os.Exit(1)
	}
	return data
}

func readDB(filePath string, reader DBReader) Recipes {
	res, err := reader.Read(filePath)
	if err != nil {
		fmt.Println("Error reading")
		os.Exit(1)
	}
	return res
}

func printJson(r Recipes) {
	data, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		fmt.Println("Fatal error")
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func printXml(r Recipes) {
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
		r := readDB(filePath, new(XmlReader))
		printJson(r)
	case "json":
		r := readDB(filePath, new(JsonReader))
		printXml(r)
	}
}
