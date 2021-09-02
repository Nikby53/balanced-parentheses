package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Nikby53/balanced-parentheses/brackets"
)

const requestURL = "http://localhost:8081/generate?n="

func request(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func workerPool(url string, cycles <-chan int, balanced chan<- bool) {
	for range cycles {
		result, err := request(url)
		if err != nil {
			log.Println(err)
		}
		balanced <- brackets.IsBalanced(result)
	}
}

func main() {
	result := make(chan bool, 1000)
	for length := 2; length <= 8; length *= 2 {
		url := requestURL + strconv.Itoa(length)

		request := make(chan int, 20)

		for i := 0; i < 20; i++ {
			go workerPool(url, request, result)
		}
		for i := 0; i < 1000; i++ {
			request <- i
		}
		close(request)
		var balanced int
		for i := 0; i < 1000; i++ {
			if <-result {
				balanced++
			}
		}
		percentOfBalanced := float64(balanced) / 1000 * 100
		fmt.Printf("lenght %v, percentage of balanced: %.1f%%\n", length, percentOfBalanced)
	}
}
