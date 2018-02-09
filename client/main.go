package main

import (
	"context"
	"fmt"
	"github.com/gitsen/playground/protos"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	server := "localhost:8080"
	conn, err := grpc.Dial(server, opts...)
	if err != nil {
		fmt.Printf("\n Error connecting %+v", err)
	}
	c := chat.NewChatClient(conn)
	stream, err := c.Talk(context.Background())
	if err != nil {
		fmt.Printf("\n Invalid response %+v", err)
	}
	stream.Send(&chat.ChatRequest{Message: "Hi"})
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s", in.Message)
			stream.Send(&chat.ChatRequest{Message: "Hi"})
		}
	}()

	<-waitc
}
