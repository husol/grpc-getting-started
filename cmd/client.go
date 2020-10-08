package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-getting-started/proto"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewHusServiceClient(conn)

	responseMessage, err := c.GetMessage(context.Background(), &proto.MessageRequest{Text: "abc"})
	if err != nil {
		log.Fatalf("Error when calling CallServer: %s", err)
	}
	log.Printf("Response from server: %s", responseMessage)

	responseAccount, err := c.GetAccount(context.Background(), &proto.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Error when calling CallServer: %s", err)
	}
	log.Printf("Response from server: %s", responseAccount)
}
