package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal("Not find index.html")
	}
	w.Write(t)
}

func main() {

	http.HandleFunc("/index.html", indexHandler)
	http.ListenAndServe(":8080", nil)

}