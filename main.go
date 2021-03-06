package main

import (
	"bootcamp/adapters"
	"bootcamp/infra/net"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	repository := adapters.NewCharacterRepository()
	getCharacterController := adapters.NewGetCharacterController(repository)
	getAllCharactersController := adapters.NewGetAllCharactersController(repository)

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Go bootcamp challenge v1")
	})

	characterRouter := router.PathPrefix("/character").Subrouter()
	characterRouter.HandleFunc("/", net.NewGetAllCharactersHandler(getAllCharactersController))
	characterRouter.HandleFunc("/{id:[0-9]+}", net.NewGetCharacterHandler(getCharacterController))

	fmt.Println("Server started")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
