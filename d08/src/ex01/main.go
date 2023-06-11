package main

import (
	"fmt"
	reflect2 "reflect"
)

type Plant interface {
	PrintAllFields()
}

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func (up *UnknownPlant) PrintAllFields() {
	PrintAllFieldsReflect(up)
}

func (aup *AnotherUnknownPlant) PrintAllFields() {
	PrintAllFieldsReflect(aup)
}

func describePlant(p Plant) {
	p.PrintAllFields()
}

func PrintAllFieldsReflect(p Plant) {
	v := reflect2.ValueOf(p).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Print(t.Field(i).Name)
		if t.Field(i).Tag != "" {
			fmt.Print("(", t.Field(i).Tag, ")")
		}
		fmt.Print(": ", v.Field(i), "\n")
	}
}

func main() {
	f := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType:    "GoodLeaf",
		Height:      15,
	}
	describePlant(&f)

	s := UnknownPlant{
		FlowerType: "SuperPlant",
		LeafType:   "SuperLeaf",
		Color:      999,
	}
	describePlant(&s)
}
