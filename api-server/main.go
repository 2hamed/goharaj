package main

import (
	"log"
	"context"

	pb "haraj/api"
	"net"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.Value)
	return &pb.StringMessage{Value: "Hello " + in.Value}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
