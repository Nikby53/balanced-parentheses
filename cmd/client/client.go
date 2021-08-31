package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const requestURL = "http://localhost:8081/generate?n=5"

func main() {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	timeout := 5 * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
