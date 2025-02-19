package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "internship202502/m/testgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedTestServer
}

func (s *server) Greet(_ context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", in.Message)
	return &pb.GreetResponse{Message: "Hello " + in.GetMessage()}, nil
}

func main() {
	flag.Parse()

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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTestServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
