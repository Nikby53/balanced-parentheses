package main

import (
	"log"
	"net/http"

	generator2 "github.com/Nikby53/balanced-parentheses/generator"

	"github.com/Nikby53/balanced-parentheses/handler"
)

func main() {
	h := handler.New(generator2.Generator{})
	http.HandleFunc("/generate", h.GenerationHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
