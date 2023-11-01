package main

import (
	"context"
	"log"
	"net"
	simple_gprc "simple-gprc/proto"

	"google.golang.org/grpc"
)

type server struct {
	simple_gprc.UnimplementedClassRegistrationServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetHelloStat(ctx context.Context, in *simple_gprc.ClassRegistrationStatRequest) (*simple_gprc.ClassRegistrationStatResponse, error) {
	return &simple_gprc.ClassRegistrationStatResponse{Result: map[int32]int32{1: 3, 2: 4}}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	simple_gprc.RegisterClassRegistrationServiceServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8081")
	log.Fatal(s.Serve(lis))
}
