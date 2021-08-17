package main

import (
	"github.com/Nikby53/balanced-parentheses/bracketsGenerator/generator"
	"log"
	"net/http"

	"github.com/Nikby53/balanced-parentheses/handler"
)

func main() {
	h := handler.New(generator.Store{})
	http.HandleFunc("/generate", h.GenerationHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
