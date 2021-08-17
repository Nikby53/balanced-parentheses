package main

import (
	"log"
	"net/http"

	"github.com/Nikby53/balanced-parentheses/handler"
)

func main() {
	http.HandleFunc("/generate", handler.GenerationHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
