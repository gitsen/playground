package server

import (
	"fmt"
	"github.com/gitsen/playground/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
)

const clientheader = "x-gitsen-client-header"

type ChatServer struct {
}

func getClientId(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if hdr, ok := md[clientheader]; ok {
			return hdr[0]
		}
	}
	return ""
}

func (c *ChatServer) Echo(stream Echo.Echo_EchoServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		stream.Send(&Echo.EchoResponse{Message: fmt.Sprintf("Hi %s Echoing %s", getClientId(stream.Context()), req.Message)})
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
	Echo.RegisterEchoServer(srv, c)
	srv.Serve(lis)
}
