package handler

import (
	"errors"
	"net/http"
	"strconv"
)

var errIncorrectInput = errors.New("incorrect input, please input a number")

var errUnexpectedServer = errors.New("unexpected server error")

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
	number, err := strconv.Atoi(req.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, errIncorrectInput.Error(), http.StatusBadRequest)
		return
	}
	temp, err := h.Generator.Generate(number)
	if err != nil {
		http.Error(w, errUnexpectedServer.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write([]byte(temp))
	if err != nil {
		http.Error(w, errUnexpectedServer.Error(), http.StatusInternalServerError)
	}
}
