package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type CandyRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

func parseArgs() CandyRequest {
	if len(os.Args[1:]) != 6 {
		log.Fatalln("Error: Wrong number of arguments")
	}

	candyType := flag.String("k", "", "accepts two-letter abbreviation for the candy type")
	countOfCandy := flag.Int("c", -1, "count of candy to buy")
	amountOfMoney := flag.Int("m", -1, "amount of money you \"gave to machine\"")
	flag.Parse()

	if *candyType == "" || *countOfCandy == -1 || *amountOfMoney == -1 {
		log.Fatalln("Error: Wrong arguments")
	}
	cr := CandyRequest{
		CandyType:  *candyType,
		CandyCount: *countOfCandy,
		Money:      *amountOfMoney,
	}
	return cr
}

func main() {
	caCert, err := ioutil.ReadFile("minica.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Создание конфигурации TLS
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	cr := parseArgs()
	mb, err := json.Marshal(cr)
	if err != nil {
		log.Fatalln(err)
		return
	}

	body := bytes.NewBuffer(mb)
	resp, err := client.Post("https://candy.tld:3334/buy_candy", "application/json", body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, _ := io.ReadAll(resp.Body)

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(b))
}
