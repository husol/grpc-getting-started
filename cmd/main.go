package main

import (
	"fmt"
	"grpc-getting-started/husgrpc"
	"grpc-getting-started/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := husgrpc.HusService{}

	grpcServer := grpc.NewServer()

	proto.RegisterHusServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}