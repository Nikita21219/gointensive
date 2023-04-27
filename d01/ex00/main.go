package main

import (
	"encoding/json"
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
	Name  string `json:"ingredient_name"`
	Count string `json:"ingredient_count"`
	Unit  string `json:"ingredient_unit"`
}

type cake struct {
	Name        string       `json:"name"`
	Time        string       `json:"time"`
	Ingredients []ingredient `json:"ingredients"`
}

type obj struct {
	Cake []cake `json:"cake"`
}

func readJson() {
	b := []byte(`{
    "cake": [
      {
        "name": "Red Velvet Strawberry Cake",
        "time": "45 min",
        "ingredients": [
          {
            "ingredient_name": "Flour",
            "ingredient_count": "2",
            "ingredient_unit": "mugs"
          },
          {
            "ingredient_name": "Strawberries",
            "ingredient_count": "7"
          },
          {
            "ingredient_name": "Vanilla extract",
            "ingredient_count": "2.5",
            "ingredient_unit": "tablespoons"
          }
        ]
      },
      {
        "name": "Blueberry Muffin Cake",
        "time": "30 min",
        "ingredients": [
          {
            "ingredient_name": "Brown sugar",
            "ingredient_count": "1",
            "ingredient_unit": "mug"
          },
          {
            "ingredient_name": "Blueberries",
            "ingredient_count": "1",
            "ingredient_unit": "mug"
          }
        ]
      }
    ]
	}`)
	var o obj

	if !json.Valid(b) {
		fmt.Println("NOT VALID!!!!!!!")
		return // TODO error?
	}
	err := json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println("Error!")
	}
	fmt.Println(o)
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
		readJson()
	}
}
