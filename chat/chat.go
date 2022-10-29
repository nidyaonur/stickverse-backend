package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message: %s", message.Body)

	return &Message{Body: "Hello " + message.Body}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}
