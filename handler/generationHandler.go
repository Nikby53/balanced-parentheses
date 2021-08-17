package handler

import (
	"net/http"
	"strconv"

	"github.com/Nikby53/balanced-parentheses/bracketsGenerator"
	"github.com/Nikby53/balanced-parentheses/bracketsGenerator/generator"
)

type Handler struct {
	Generator generator.GenerationRepo
}

func New(gen generator.GenerationRepo) Handler {
	return Handler{Generator: gen}
}

func (h Handler) GenerationHandler(w http.ResponseWriter, req *http.Request) {
	number, err := strconv.Atoi(req.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, bracketsGenerator.ErrIncorrectInput.Error(), http.StatusBadRequest)
		return
	}
	temp, _ := h.Generator.Generation(number)
	_, err = w.Write([]byte(temp))
	if err != nil {
		http.Error(w, bracketsGenerator.ErrIncorrectInput.Error(), http.StatusBadRequest)
	}
}
