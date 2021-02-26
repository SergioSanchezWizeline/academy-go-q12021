package net

import (
	"bootcamp/adapters"
	"bootcamp/helper"
	"fmt"
	"net/http"
)

func NewGetCharacterHandler(getCharacterController adapters.GetCharacterController) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := helper.ExtractId("/character/", request.URL.Path)
		if err != nil {
			fmt.Fprintf(writer, "{\"error\": \"%v\"}", err)
			return
		}
		result, err := getCharacterController.Execute(id)
		if err != nil {
			fmt.Fprintf(writer, "{\"error\": \"%v\"}", err)
			return
		}

		fmt.Fprint(writer, result)
	}
}
