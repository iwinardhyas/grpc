package main

import (
	"context"
	"log"

	chat "grpc/chat"

	"google.golang.org/grpc"
)

func main() {
	var con *grpc.ClientConn
	con, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %S", err)
	}
	defer con.Close()

	c := chat.NewChatserviceClient(con)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "hello from client"})
	if err != nil {
		log.Fatalf("error when calling sayhello: %s", err)
	}
	log.Printf("response from server: %s", response.Body)
}
