package handler

import (
	"net/http"
	"strconv"

	"github.com/Nikby53/balanced-parentheses/bracketsGenerator"
	"github.com/Nikby53/balanced-parentheses/bracketsGenerator/generator"
)

func GenerationHandler(w http.ResponseWriter, req *http.Request) {
	number, err := strconv.Atoi(req.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, bracketsGenerator.ErrIncorrectInput.Error(), http.StatusBadRequest)
		return
	}
	temp, _ := generator.Store{}.Generation(number)
	_, err = w.Write([]byte(temp))
	if err != nil {
		http.Error(w, bracketsGenerator.ErrIncorrectInput.Error(), http.StatusBadRequest)
	}
}
