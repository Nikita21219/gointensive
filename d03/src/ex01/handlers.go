package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Place struct {
	Name    string
	Address string
	Phone   string
}

func renderTemplate(page string, w http.ResponseWriter, data any) {
	files := []string{
		page + ".page.html",
		"base.layout.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	indexName := "places"

	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	dataErr := map[string]string{
		"Error": fmt.Sprintf("Invalid 'page' value: '%s'\n", pageParam),
	}
	if err != nil || page <= 0 {
		renderTemplate("index", w, dataErr)
		return
	}

	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body: bytes.NewReader([]byte(`{
				"size": 20000,
				"query": {
					"match_all": {}
				}
			}`)),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	places := make([]Place, 0, 20000)
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		place := Place{
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
		}
		places = append(places, place)
	}

	total := len(places)
	limit := 10
	offset := (page - 1) * limit

	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		renderTemplate("index", w, dataErr)
		return
	}

	prevPage := -1
	if page > 1 {
		prevPage = page - 1
	}

	nextPage := -1
	if page < lastPage {
		nextPage = page + 1
	}

	data := map[string]interface{}{
		"Places":   places[offset : offset+limit],
		"Total":    total,
		"PrevPage": prevPage,
		"NextPage": nextPage,
		"LastPage": lastPage,
	}

	renderTemplate("index", w, data)
}
