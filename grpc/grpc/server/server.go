package main

import (
	"context"
	pb "demo-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type Server struct {
	pb.UnsafeHelloServiceServer
}

func (s *Server) Hello(ctx context.Context, in *pb.String) (*pb.String, error) {
	return in, nil
}

func main() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
