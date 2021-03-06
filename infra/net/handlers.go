package net

import (
	"bootcamp/adapters"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewGetCharacterHandler(getCharacterController adapters.GetCharacterController) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
			return
		}
		result, err := getCharacterController.Execute(id)
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
			return
		}

		fmt.Fprint(w, result)
	}
}

func NewGetAllCharactersHandler(getCharacterController adapters.GetAllCharactersController) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := getCharacterController.Execute()
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%v\"}", err)
			return
		}

		fmt.Fprint(w, result)
	}
}
