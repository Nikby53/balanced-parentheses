package main

import (
	"log"
	"net/http"

	"github.com/Nikby53/balanced-parentheses/generator"
	"github.com/Nikby53/balanced-parentheses/handler"
)

func main() {
	h := handler.New(generator.Generator{})
	http.HandleFunc("/generate", h.GenerationHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
