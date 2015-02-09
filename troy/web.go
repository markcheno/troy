package main

import (
	"encoding/json"
	"github.com/ironbay/troy"
	"io/ioutil"
	"log"
	"net/http"
)

func web() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		obj := troy.Script(string(body))
		log.Println(string(body))
		w.Header().Set("Content-Type", "application/json")
		json, _ := json.Marshal(obj)
		w.Write(json)
	})
	port := ":9876"
	log.Println("Listening on", port)
	http.ListenAndServe(port, nil)
}
