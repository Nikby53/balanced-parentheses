package handler

import (
	"errors"
	"net/http"
	"strconv"
)

// Handler is struct that holds BracketsGenerator
// interface and responsible for all possible handlers which are its methods.
type Handler struct {
	Generator BracketsGenerator
}

var errIncorrectInput = errors.New("incorrect input, please input a number")

// New function is a constructor for Handler.
func New(b BracketsGenerator) Handler {
	return Handler{Generator: b}
}

// BracketsGenerator is an interface that describes method
// for generating parentheses.
type BracketsGenerator interface {
	Generate(num int) (string, error)
}

// GenerationHandler is a handler.
func (h Handler) GenerationHandler(w http.ResponseWriter, req *http.Request) {
	number, err := strconv.Atoi(req.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, errIncorrectInput.Error(), http.StatusBadRequest)
		return
	}
	temp, _ := h.Generator.Generate(number)
	_, err = w.Write([]byte(temp))
	if err != nil {
		http.Error(w, errIncorrectInput.Error(), http.StatusBadRequest)
	}
}
