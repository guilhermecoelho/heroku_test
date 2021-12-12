package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Moviment struct {
	Id    int     `json:"id"`
	Value float64 `json:"value"`
}

func main() {

	routes := mux.NewRouter().StrictSlash(true)

	getRouter := routes.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/", test)

	log.Fatal(http.ListenAndServe(":8080", routes))

}

func test(resp http.ResponseWriter, r *http.Request) {

	mov := Moviment{1, 35}

	e := json.NewEncoder(resp)
	e.Encode(mov)
}
