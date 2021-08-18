package handler

import (
	"errors"
	"net/http"
	"strconv"
)

type Handler struct {
	Generator GenerationRepo
}

// ErrIncorrectInput checks for incorrect input.
var ErrIncorrectInput = errors.New("incorrect input, please input a number")

func New(gen GenerationRepo) Handler {
	return Handler{Generator: gen}
}

type GenerationRepo interface {
	Generation(num int) (string, error)
}

func (h Handler) GenerationHandler(w http.ResponseWriter, req *http.Request) {
	number, err := strconv.Atoi(req.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, ErrIncorrectInput.Error(), http.StatusBadRequest)
		return
	}
	temp, _ := h.Generator.Generation(number)
	_, err = w.Write([]byte(temp))
	if err != nil {
		http.Error(w, ErrIncorrectInput.Error(), http.StatusBadRequest)
	}
}
