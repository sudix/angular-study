package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := []map[string]interface{} {
		{"name": "hoge", "age": 10},
		{"name": "fuga", "age": 30},
		{"name": "piyo", "age": 60},
	}

	// wait few seconds to test async loading
	time.Sleep(2 * time.Second)

	b, _ := json.Marshal(list)
	fmt.Fprint(w, string(b))
}

func main() {
	// json response
	http.HandleFunc("/json", jsonHandler)

	// static file serve
	hs := http.FileServer(http.Dir("./public"))
	http.Handle("/", hs) //http.StripPrefix("/public/", hs))

	port := 8080
	log.Printf("Listening on %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
