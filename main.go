package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var (
	cities = []string {"San Francisco, Amsterdam, Berlin, New York"}
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(cities)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
