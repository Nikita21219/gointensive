package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
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

func getRecipies(filePath string) Recipes {
	var r Recipes
	switch strings.Split(filePath, ".")[1] {
	case "xml":
		r = readDB(filePath, new(XmlReader))
	case "json":
		r = readDB(filePath, new(JsonReader))
	}
	return r
}

func checkOnlyNewIngredients(ingrs []Ingredient, cake string, diff *[]string) {
	for _, ingr := range ingrs {
		*diff = append(
			*diff,
			fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingr.Name, cake))
	}
}

func checkIngredientCountAndUnit(ingr1 Ingredient, ingr2 Ingredient, diff *[]string, cake string) {
	if ingr1.Unit != ingr2.Unit {
		if ingr1.Unit != "" && ingr2.Unit == "" {
			*diff = append(
				*diff,
				fmt.Sprintf(
					"REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
					ingr1.Unit,
					ingr1.Name,
					cake))
		} else if ingr1.Unit == "" && ingr2.Unit != "" {
			*diff = append(
				*diff,
				fmt.Sprintf(
					"ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
					ingr2.Unit,
					ingr2.Name,
					cake))
		} else {
			*diff = append(
				*diff,
				fmt.Sprintf(
					"CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
					ingr1.Name,
					cake,
					ingr2.Unit,
					ingr1.Unit))
		}
	}
	if ingr1.Count != ingr2.Count {
		*diff = append(
			*diff,
			fmt.Sprintf(
				"CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
				ingr1.Name,
				cake,
				ingr2.Count,
				ingr1.Count))
	}
}

func checkIngredientsDiff(
	ingr1 []Ingredient,
	ingr2 []Ingredient,
	cakeName1 string,
	cakeName2 string,
	diff *[]string) {
	for i, el := range ingr1 {
		if len(ingr2) >= i+1 {
			if el.Name != ingr2[i].Name {
				*diff = append(
					*diff,
					fmt.Sprintf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingr2[i].Name, cakeName1))
				*diff = append(
					*diff,
					fmt.Sprintf("REMOVED ingredient \"%s\" for cake \"%s\"\n", el.Name, cakeName2))
			} else {
				checkIngredientCountAndUnit(el, ingr2[i], diff, cakeName1)
			}
		}
	}

	if len(ingr1) > len(ingr2) {
		checkOnlyNewIngredients(ingr1[len(ingr2):], cakeName1, diff)
	} else if len(ingr1) < len(ingr2) {
		checkOnlyNewIngredients(ingr2[len(ingr1):], cakeName2, diff)
	}
}

func getDiffCake(cake1 Cake, cake2 Cake) ([]string, []string, []string) {
	addedAndRemoved := make([]string, 0)
	changed := make([]string, 0)
	ingredientsDiff := make([]string, 0)

	if cake1.Name != cake2.Name {
		addedAndRemoved = append(addedAndRemoved, fmt.Sprintf("ADDED cake \"%s\"\n", cake2.Name))
		addedAndRemoved = append(addedAndRemoved, fmt.Sprintf("REMOVED cake \"%s\"\n", cake1.Name))
	} else {
		if cake1.Time != cake2.Time {
			str := fmt.Sprintf(
				"CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
				cake2.Name,
				cake2.Time,
				cake1.Time)
			changed = append(changed, str)
		}
		checkIngredientsDiff(cake1.Ingredients, cake2.Ingredients, cake1.Name, cake2.Name, &ingredientsDiff)
	}

	return addedAndRemoved, changed, ingredientsDiff
}

func printArr(arr []string) {
	for _, str := range arr {
		fmt.Print(str)
	}
}

//func printDiffCakes(r1 []Cake, r2 []Cake) {
//	addedAndRemoved := make([]string, 0)
//	changed := make([]string, 0)
//	ingredientsDiff := make([]string, 0)
//
//	for idx, cake := range r1 {
//		if len(r2) >= idx+1 {
//			arr1, arr2, arr3 := getDiffCake(cake, r2[idx])
//			addedAndRemoved = append(addedAndRemoved, arr1...)
//			changed = append(changed, arr2...)
//			ingredientsDiff = append(ingredientsDiff, arr3...)
//		}
//	}
//
//	printArr(addedAndRemoved)
//	for i := len(r1); i < len(r2); i++ {
//		fmt.Printf("ADDED cake \"%s\"\n", r2[i].Name)
//	}
//	printArr(changed)
//	printArr(ingredientsDiff)
//}

func searchByName(r []Cake, name string) int {
	for idx := range r {
		if r[idx].Name == name {
			return idx
		}
	}
	return -1
}

func removeByIdx(a []Cake, i int) []Cake {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = Cake{}
	a = a[:len(a)-1]
	return a
}

func printDiffCakes(r1 []Cake, r2 []Cake) {
	addedAndRemoved := make([]string, 0)
	changed := make([]string, 0)
	ingredientsDiff := make([]string, 0)

	for _, cake := range r1 {
		idx := searchByName(r2, cake.Name)
		if idx == -1 {
			addedAndRemoved = append(addedAndRemoved, fmt.Sprintf("REMOVED cake \"%s\"\n", cake.Name))
		} else {
			arr1, arr2, arr3 := getDiffCake(cake, r2[idx])
			addedAndRemoved = append(addedAndRemoved, arr1...)
			changed = append(changed, arr2...)
			ingredientsDiff = append(ingredientsDiff, arr3...)
			r2 = removeByIdx(r2, idx)
		}
	}

	printArr(addedAndRemoved)
	for _, cake := range r2 {
		str := fmt.Sprintf("ADDED cake \"%s\"\n", cake.Name)
		fmt.Printf(str)
	}
	printArr(changed)
	printArr(ingredientsDiff)
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

	rOld := getRecipies(oldFile).Cake
	rNew := getRecipies(newFile).Cake

	printDiffCakes(rOld, rNew)
}
