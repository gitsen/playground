package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gitsen/playground/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"time"
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
	c := Echo.NewEchoClient(conn)
	ctx := context.Background()
	md := metadata.New(map[string]string{clientheader: *clientId})
	ctx = metadata.NewOutgoingContext(ctx, md)

	stream, err := c.Echo(ctx)
	if err != nil {
		fmt.Printf("\n Invalid response %+v", err)
	}
	stream.Send(&Echo.EchoRequest{Message: "Hi"})
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
			time.Sleep(time.Second * 5)
			stream.Send(&Echo.EchoRequest{Message: "Hi"})
		}
	}()

	<-waitc
}
