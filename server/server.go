package server

import (
	"fmt"
	"github.com/gitsen/playground/protos"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type ChatServer struct {
}

func (c *ChatServer) Talk(stream chat.Chat_TalkServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		stream.Send(&chat.ChatResponse{Message: fmt.Sprintf("Echoing %s", req.Message)})
	}
	<-stream.Context().Done()
	return stream.Context().Err()
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
