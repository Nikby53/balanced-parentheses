package main

import (
	"log"
	"net/http"

	"github.com/Nikby53/balanced-parentheses/brackets"

	"github.com/Nikby53/balanced-parentheses/handler"
)

func main() {
	h := handler.New(brackets.Generator{})
	http.HandleFunc("/generate", h.GenerationHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
