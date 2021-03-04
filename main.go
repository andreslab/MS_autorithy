package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//devuelve la ruta
type route struct {
	Url string `json:url`
}

//cifra el request
type reqEncode struct {
	Encode string `json:encode`
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my api")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":3000", router))
}
