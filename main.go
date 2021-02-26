package main

import (
	"bootcamp/adapters"
	"bootcamp/infra/net"
	"fmt"
	"log"
	"net/http"
)

func main() {

	repository := adapters.NewCharacterRepository()
	getCharacterController := adapters.NewGetCharacterController(repository)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Go bootcamp challenge v1")
	})

	http.HandleFunc("/character/", net.NewGetCharacterHandler(getCharacterController))

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
