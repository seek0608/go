package main

import (
	"context"
	pb "demo-grpcgateway/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s server) Hello(ctx context.Context, in *pb.String) (*pb.String, error) {
	return in, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	err = s.Serve(listen)
	if err = s.Serve(listen); err != nil {
		panic(err)
	}
}
