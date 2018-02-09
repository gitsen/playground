package server

import (
	"context"
	"fmt"
	"github.com/gitsen/playground/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ChatServer struct {
}

func (c *ChatServer) Talk(ctx context.Context, req *chat.ChatRequest) (*chat.ChatResponse, error) {
	return &chat.ChatResponse{Message: "Hi There"}, nil
}

func (c *ChatServer) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	chat.RegisterChatServer(srv, c)
	srv.Serve(lis)
}
