package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("Error starting TCP server: %s", err)
	}
	defer listener.Close()

	log.Println("TCP server listening on :80")
	for {
		_, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}

	}
}
