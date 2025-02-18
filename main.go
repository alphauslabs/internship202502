package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[healthcheck] Received request")
		w.WriteHeader(http.StatusOK)
	})
	log.Println("Serving at :80")
	http.ListenAndServe(":80", nil)
}
