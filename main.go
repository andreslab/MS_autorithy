package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type reqRoute struct {
	Tag string `json:tag`
}

//devuelve la ruta
type resRoute struct {
	Url string `json:url`
}

//cifra el request
type reqEncode struct {
	Value string `json:value`
}

type resEncode struct {
	Value string `json:value`
}

//decifra el request
type reqDecode struct {
	Value string `json:value`
}

type resDecode struct {
	Value string `json:value`
}

type Route struct {
	Tag   string `json:Tag`
	Value string `json:Value`
}

func getRoute(w http.ResponseWriter, r *http.Request) {

	var dataReqRoute reqRoute
	var dataResRoute resRoute
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &dataReqRoute)

	routes, err := getRoutesFromJson()
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	url, err := searchRouteFromTag("registro", routes)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	dataResRoute = resRoute{
		Url: url,
	}

	json.NewEncoder(w).Encode(dataResRoute)
}
func getEncode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my api")
}
func getDecode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my api")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/route", getRoute)
	router.HandleFunc("/encode", getEncode)
	router.HandleFunc("/decode", getDecode)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getRoutesFromJson() ([]Route, error) {
	var routes = []Route{
		{
			Tag:   "registro",
			Value: "localhost:4000/api/register2",
		},
		{
			Tag:   "login",
			Value: "localhost:4000/api/register",
		},
		{
			Tag:   "token",
			Value: "localhost:4000/api/register",
		},
	}

	/*data, err := json.Marshal(routes)
	if err != nil {
		log.Fatal("Error al convertir a JSON" + err.Error())
	}
	fmt.Printf("%s", data)*/

	return routes, nil

}

func searchRouteFromTag(tag string, routes []Route) (string, error) {
	for _, r := range routes {
		if r.Tag == tag {
			return r.Value, nil
		}
	}
	return "", errors.New("Not found")
}
