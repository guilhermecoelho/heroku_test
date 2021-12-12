package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Moviment struct {
	Id    int     `json:"id"`
	Value float64 `json:"value"`
}

func main() {

	godotenv.Load()

	routes := mux.NewRouter().StrictSlash(true)

	getRouter := routes.Methods(http.MethodGet).Subrouter()

	getRouter.HandleFunc("/", test)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), routes))

}

func test(resp http.ResponseWriter, r *http.Request) {

	mov := Moviment{1, 35}

	e := json.NewEncoder(resp)
	e.Encode(mov)
}
