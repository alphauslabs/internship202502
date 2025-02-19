package main

import (
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received!")
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    ":50051",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	log.Println("serving http on :50051")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting HTTP server: %s", err)
	}

}
