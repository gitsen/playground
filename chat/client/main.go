package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/gitsen/playground/chat/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
)

const clientheader = "x-gitsen-client-header"

func main() {
	clientId := flag.String("clientId", "", "Identifier of the client")
	flag.Parse()
	if *clientId == "" {
		panic("Client id is needed")
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	server := "localhost:8080"
	conn, err := grpc.Dial(server, opts...)
	if err != nil {
		fmt.Printf("\n Error connecting %+v", err)
	}
	c := Chat.NewChatClient(conn)
	ctx := context.Background()
	md := metadata.New(map[string]string{clientheader: *clientId})
	ctx = metadata.NewOutgoingContext(ctx, md)
	c.Register(ctx, &Chat.RegisterRequest{ClientId: *clientId})
	waitc := make(chan struct{})
	go chat(c, ctx, waitc)
	<-waitc
}

func send(stream Chat.Chat_BroadcastClient) {
	fmt.Println("Happy chatting...")
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if sc.Scan() {
			stream.Send(&Chat.BroadcastRequest{Message: sc.Text()})
		} else {
			panic(fmt.Sprintf("Error reading from terminal %+v", sc.Err()))
		}
	}
}

func chat(c Chat.ChatClient, ctx context.Context, waitc chan struct{}) {
	stream, err := c.Broadcast(ctx)
	if err != nil {
		fmt.Printf("\n Bad response, %+v", err)
		close(waitc)
		return
	}
	go send(stream)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			close(waitc)
			return
		}
		if err != nil {
			log.Fatalf("Receipt failed : %v", err)
		}
		fmt.Printf("\n> %s\n> ", resp.Message)
	}
}
