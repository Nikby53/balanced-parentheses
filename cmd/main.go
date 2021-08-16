package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Nikby53/balanced-parentheses/bracketsGenerator"

	"github.com/Nikby53/balanced-parentheses/bracketsGenerator/generator"
)

func generationHandler(w http.ResponseWriter, req *http.Request) {
	number, err := strconv.Atoi(req.URL.Query().Get("n"))
	if err != nil {
		fmt.Fprintf(w, "%v", bracketsGenerator.ErrIncorrectInput)
	}
	temp, _ := generator.Generation(number)
	_, err = w.Write([]byte(temp))
	if err != nil {
		fmt.Fprintf(w, "%v", bracketsGenerator.ErrIncorrectInput)
	}

}

func main() {
	http.HandleFunc("/generate", generationHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
