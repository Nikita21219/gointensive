package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Place struct {
	Id            float64  `json:"id"`
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Phone         string   `json:"phone"`
	LocationField Location `json:"location"`
}

type ResponsePlaces struct {
	Name     string  `json:"name"`
	Total    int     `json:"total"`
	Places   []Place `json:"places"`
	PrevPage int     `json:"prev_page"`
	NextPage int     `json:"next_page"`
	LastPage int     `json:"last_page"`
}

func responseAPI(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write(data)
	if err != nil {
		fmt.Println("Error write data:", err)
	}
}

func getHits(indexName string) ([]interface{}, error) {
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
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("HTTP error")
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	return hits, nil
}

func getPlaces(hits []interface{}) []Place {
	places := make([]Place, 0, 20000)
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		location := source["location"].(map[string]interface{})
		place := Place{
			Id:      source["id"].(float64),
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
			LocationField: Location{
				location["lat"].(float64),
				location["lon"].(float64),
			},
		}
		places = append(places, place)
	}
	return places
}

func index(w http.ResponseWriter, r *http.Request) {
	indexName := "places"

	// Get page param
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	dataErr, err := json.Marshal(map[string]string{
		"error": fmt.Sprintf("Invalid 'page' value: '%s'", pageParam),
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Handle error page param
	if page <= 0 {
		responseAPI(w, dataErr)
		return
	}

	// Get hits from elastic
	hits, err := getHits(indexName)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Get Places from hits
	places := getPlaces(hits)

	total := len(places)
	limit := 10
	offset := (page - 1) * limit

	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		responseAPI(w, dataErr)
		return
	}

	rp := ResponsePlaces{
		Name:     "Places",
		Total:    total,
		Places:   places[offset : offset+limit],
		LastPage: lastPage,
	}
	if page > 1 {
		rp.PrevPage = page - 1
	}
	if page < lastPage {
		rp.NextPage = page + 1
	}

	data, err := json.Marshal(rp)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	responseAPI(w, data)
}
