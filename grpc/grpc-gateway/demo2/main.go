package main

import (
	"context"
	gw "demo-grpcgateway2/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func main() {
	run()
}

func run() error {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterMessageServiceHandlerFromEndpoint(context.Background(), mux, ":8080", opts)
	if err != nil {
		panic(err)
	}
	return http.ListenAndServe(":8081", mux)
}
