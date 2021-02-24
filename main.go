package main

import (
	"bootcamp/adapters"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	repository := adapters.NewCharacterRepository()
	getCharacterController := adapters.NewGetCharacterController(repository)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello there!")
	})

	http.HandleFunc("/character/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/character/"):]
		id, err := strconv.Atoi(idStr)
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
	})

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
