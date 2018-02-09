package main

import (
	"context"
	"fmt"
	"github.com/gitsen/playground/protos"
	"google.golang.org/grpc"
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
	resp, err := c.Talk(context.Background(), &chat.ChatRequest{Message: "Hello World"})
	if err != nil {
		fmt.Printf("\n Invalid response %+v", err)
	}
	fmt.Println(resp.Message)
}
