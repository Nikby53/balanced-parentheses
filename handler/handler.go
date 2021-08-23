package handler

import (
	"errors"
	"net/http"
	"strconv"
)

type generationRequest struct {
	number int
}

var (
	errNumberRequired = errors.New("number query parameters is required")
	errShouldBeNumber = errors.New("query parameter should be number")
	errLessThanZero   = errors.New("parameter should be greater than zero")
)

// Validate method is for validating handler request.
func (g *generationRequest) Validate(r *http.Request) error {
	query := r.URL.Query()
	number := query.Get("n")
	if number == "" {
		return errNumberRequired
	}
	temp, err := strconv.Atoi(number)
	if err != nil {
		return errShouldBeNumber
	}
	if temp < 1 {
		return errLessThanZero
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
	var genHandler generationRequest
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
