package main

//extern char *ask_cow(char phrase[]);
import "C"

import (
	"encoding/json"
	"log"
	"net/http"
)

type CandyRequest struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type CandyResponse struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var candyPrices = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func checkRequest(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != "/buy_candy" {
		http.NotFound(w, r)
		return true
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return true
	}
	return false
}

func notFound(t string) bool {
	_, exists := candyPrices[t]
	return !exists
}

func notValidData(cr CandyRequest, w http.ResponseWriter) bool {
	if cr.CandyCount < 0 || cr.Money < 0 || notFound(cr.CandyType) {
		er := ErrorResponse{Error: "some error in input data"}
		b, err := json.Marshal(er)

		if err != nil {
			http.Error(w, "Error Marshal", 500)
			return true
		}

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(b)
		if err != nil {
			http.Error(w, "Error write bytes", 500)
			return true
		}

		return true
	}
	return false
}

func success(w http.ResponseWriter, change int) {
	str := "Thank you!"
	cow := C.ask_cow(C.CString(str))
	thanks := C.GoString(cow)

	cr := CandyResponse{Change: change, Thanks: thanks}
	b, err := json.Marshal(cr)
	if err != nil {
		http.Error(w, "Error Marshal", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Error write bytes", 500)
		return
	}
}

func fail(w http.ResponseWriter) {
	er := ErrorResponse{Error: "not enough money"}
	b, err := json.Marshal(er)

	if err != nil {
		http.Error(w, "Error Marshal", 500)
		return
	}

	w.WriteHeader(http.StatusPaymentRequired)
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Error write bytes", 500)
		return
	}
}

func buyCandy(w http.ResponseWriter, r *http.Request) {
	if checkRequest(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	candyRequest := CandyRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&candyRequest)

	if err != nil {
		log.Println("Error decode:", err)
		http.Error(w, "Error decode", 500)
		return
	}

	if notValidData(candyRequest, w) {
		return
	}

	amount := candyRequest.CandyCount * candyPrices[candyRequest.CandyType]
	if amount <= candyRequest.Money {
		success(w, candyRequest.Money-amount)
	} else {
		fail(w)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/buy_candy", buyCandy)

	err := http.ListenAndServeTLS(":3334", "../ex01/candy.tld/cert.pem", "../ex01/candy.tld/key.pem", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
