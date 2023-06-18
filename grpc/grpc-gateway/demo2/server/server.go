package main

import (
	"context"
	pb "demo-grpcgateway2/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s server) Hello(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return in, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})
	err = s.Serve(listen)
	if err = s.Serve(listen); err != nil {
		panic(err)
	}
}
