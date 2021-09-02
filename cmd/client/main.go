package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Nikby53/balanced-parentheses/brackets"
)

const (
	requestURL       = "http://localhost:8081/generate?n="
	numberOfRequests = 1000
)

// request makes a request to a URL
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

// workerPool implements the required number of cycles and writes
// the result of isBalanced function to the balanced channel.
func workerPool(url string, cycles <-chan int, balanced chan<- bool) {
	for range cycles {
		result, err := request(url)
		if err != nil {
			log.Println(err)
			balanced <- false
			continue
		}
		balanced <- brackets.IsBalanced(result)
	}
}

func main() {
	result := make(chan bool, numberOfRequests)
	for length := 2; length <= 8; length *= 2 {
		url := requestURL + strconv.Itoa(length)
		cycles := make(chan int, 20)
		for i := 0; i < 20; i++ {
			go workerPool(url, cycles, result)
		}
		for i := 0; i < numberOfRequests; i++ {
			cycles <- i
		}
		close(cycles)
		var balanced int
		for i := 0; i < numberOfRequests; i++ {
			if <-result {
				balanced++
			}
		}
		percentOfBalanced := float64(balanced) / numberOfRequests * 100
		fmt.Printf("length %v, percentage of balanced: %.1f%%\n", length, percentOfBalanced)
	}
}
