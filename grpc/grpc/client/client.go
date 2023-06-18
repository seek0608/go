package main

import (
	pb "demo-grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	response, err := c.Hello(context.Background(), &pb.String{Value: "test"})
	if err != nil {
		log.Fatalf("Error when calling QuickSort: %s", err)
	}
	log.Printf("Response from server: %v", response.Value)
}
