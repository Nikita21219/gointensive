package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
)

var secretKey = []byte("My Secret Key 42")

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
	Name   string  `json:"name"`
	Places []Place `json:"places"`
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

func getSearchResult(indexName, lat, lon string) ([]Place, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"sort": []map[string]interface{}{
			{
				"_geo_distance": map[string]interface{}{
					"location": map[string]interface{}{
						"lat": lat,
						"lon": lon,
					},
					"order":           "asc",
					"unit":            "km",
					"mode":            "min",
					"distance_type":   "arc",
					"ignore_unmapped": true,
				},
			},
		},
		"size": 3,
	}

	body, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(string(body)),
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

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Error convert")
	}

	places := make([]Place, 0, 3)

	for _, hit := range hits {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Error convert")
		}
		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Error convert")
		}
		data, err := json.Marshal(source)
		if err != nil {
			return nil, err
		}
		var place Place
		err = json.Unmarshal(data, &place)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}

func getTokenFromAuthorizationHeader(authorizationHeader string) (string, error) {
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("incorrect authorization header format")
	}

	return parts[1], nil
}

func index(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromAuthorizationHeader(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Not authorized", 401)
		return
	}

	verifyed, err := verifyJWT(token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if !verifyed {
		http.Error(w, "Not authorized", 401)
		return
	}

	indexName := "places"

	// Get lat and lon params
	latParam := r.URL.Query().Get("lat")
	lonParam := r.URL.Query().Get("lon")

	places, err := getSearchResult(indexName, latParam, lonParam)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}

	rp := ResponsePlaces{
		Name:   "Recommendation",
		Places: places,
	}

	data, err := json.Marshal(rp)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	responseAPI(w, data)
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyJWT(token string) (bool, error) {
	tokenRes, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unsupported signature algorithm: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return false, err
	}

	return tokenRes.Valid, nil
}

func getToken(w http.ResponseWriter, r *http.Request) {
	token, err := generateJWT()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	data, err := json.Marshal(map[string]interface{}{
		"Token": token,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	responseAPI(w, data)
}
