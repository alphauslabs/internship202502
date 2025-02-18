package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Error starting TCP server: %s", err)
	}
	defer listener.Close()

	log.Println("TCP server listening on :80")

	go func() {
		for {
			c, err := listener.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %s", err)
				continue
			}
			log.Println("Accepted connection for health checks")
			c.Close()
		}
	}()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received!")
		w.WriteHeader(http.StatusOK)
	})

	log.Println("serving http on :50051")
	err = http.ListenAndServe(":50051", nil)
	if err != nil {
		log.Fatalf("Error starting HTTP server: %s", err)
	}
}
