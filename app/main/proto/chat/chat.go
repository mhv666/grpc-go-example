package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) sayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recived a message body from client: %s", message.Body)
	return &Message{Body: "Hello from the Server"}, nil
}
