package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"os"
)

type DBReader interface {
	Read(filePath string) (Schema, error)
}

type JsonReader struct{}

type PropertiesType struct {
	Type string `json:"type"`
}

type Properties struct {
	Name     PropertiesType `json:"name"`
	Address  PropertiesType `json:"address"`
	Phone    PropertiesType `json:"phone"`
	Location PropertiesType `json:"location"`
}

type Schema struct {
	Properties Properties `json:"properties"`
}

func (j *JsonReader) Read(filePath string) (Schema, error) {
	b := getDataFromFile(filePath)
	var r Schema

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

func readDB(filePath string, reader DBReader) Schema {
	res, err := reader.Read(filePath)
	if err != nil {
		log.Fatalln("Error reading file.", err)
	}
	return res
}

func getDataFromFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("File does not exists")
		os.Exit(1)
	}
	return data
}

func createIndex(indexName string, es *elastic.Client) error {
	if _, err := es.Indices.Delete([]string{indexName}); err != nil {
		return err
	}

	res, err := es.Indices.Create(indexName)
	if err != nil {
		return err
	}
	if res.IsError() {
		return errors.New("error")
	}
	return nil
}

func mappingIndex(indexName string, es *elastic.Client) error {
	// Marshal json from file
	schema := readDB("schema.json", new(JsonReader))
	schemaBytes, err := json.Marshal(schema)
	if err != nil {
		return err
	}

	fmt.Println(string(schemaBytes))

	// Creatindg index request
	indexReq := esapi.IndexRequest{
		Index: indexName,
		Body:  bytes.NewReader(schemaBytes),
	}

	// Execute index request
	resp, err := indexReq.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle errors
	if resp.IsError() {
		//fmt.Println(resp.String())
		return fmt.Errorf("%s:\n%s", resp.Status(), resp.String())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}

		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			return err
		} else {
			fmt.Println(r)
		}
	}
	return nil
}

func main() {
	indexName := "places"

	// Create ES client
	es, err := elastic.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	}

	// Create index
	err = createIndex(indexName, es)
	if err != nil {
		log.Fatalf("Error creating the index: %s\n", err)
	}

	// Add mapping to index
	err = mappingIndex(indexName, es)
	if err != nil {
		log.Fatalf("Error mapping the index: %s\n", err)
	}
}
