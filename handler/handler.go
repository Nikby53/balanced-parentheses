package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

type generationHandler struct {
	number int
}

// Validate method is for validating handler.
func (g *generationHandler) Validate(r *http.Request) error {
	query := r.URL.Query()
	number := query.Get("n")
	if number == "" {
		return fmt.Errorf("number query parametres is required")
	}
	temp, err := strconv.Atoi(number)
	if err != nil {
		return fmt.Errorf("query parameter should be number")
	}
	if temp < 1 {
		return fmt.Errorf("parameter should be greater than zero")
	}
	g.number = temp
	return nil
}

// BracketsGenerator is an interface that describes method
// for generating parentheses.
type BracketsGenerator interface {
	Generate(num int) (string, error)
}

// Handler is struct that holds BracketsGenerator
// interface and responsible for all possible handlers which are its methods.
type Handler struct {
	Generator BracketsGenerator
}

// New function is a constructor for Handler.
func New(b BracketsGenerator) Handler {
	return Handler{Generator: b}
}

// GenerationHandler is a handler.
func (h Handler) GenerationHandler(w http.ResponseWriter, req *http.Request) {
	var genHandler generationHandler
	err := genHandler.Validate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	temp, err := h.Generator.Generate(genHandler.number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write([]byte(temp))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
