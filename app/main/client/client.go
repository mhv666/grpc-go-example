package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/mhv666/grpc-go-example/proto/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "hello body",
	}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when call sayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)

}
