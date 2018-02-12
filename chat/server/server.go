package server

import (
	"fmt"
	"github.com/gitsen/playground/chat/protos"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
	"sync"
)

const clientheader = "x-gitsen-client-header"

type clientBroadcast struct {
	ClientId string
	resp     Chat.BroadcastResponse
}

type ChatServer struct {
	clients           map[string]chan Chat.BroadcastResponse
	broadcastChan     chan clientBroadcast
	clientStreamsLock sync.RWMutex
}

func New() *ChatServer {
	c := &ChatServer{
		clients:       make(map[string]chan Chat.BroadcastResponse),
		broadcastChan: make(chan clientBroadcast, 100),
	}
	go c.broadcast()
	return c
}

func getClientId(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if hdr, ok := md[clientheader]; ok {
			return hdr[0]
		}
	}
	return ""
}

func (c *ChatServer) Register(ctx context.Context, req *Chat.RegisterRequest) (*Chat.RegisterResponse, error) {
	if req.ClientId == "" {
		return nil, errors.New("No client Id passed in")
	}
	c.clientStreamsLock.Lock()
	c.clients[req.ClientId] = make(chan Chat.BroadcastResponse, 100)
	c.clientStreamsLock.Unlock()
	log.Printf("Registered client %s", req.ClientId)
	return &Chat.RegisterResponse{}, nil
}

func (c *ChatServer) broadcast() {
	for {
		select {
		case res := <-c.broadcastChan:
			c.clientStreamsLock.RLock()
			for c, s := range c.clients {
				if res.ClientId == c {
					continue
				}
				log.Printf("Broadcasting to %s", c)
				s <- res.resp
			}
			c.clientStreamsLock.RUnlock()
		}
	}
}

func (c *ChatServer) Broadcast(stream Chat.Chat_BroadcastServer) error {
	clientId := getClientId(stream.Context())
	if clientId == "" {
		return errors.New("Client not registered")
	}
	go c.sendBroadcastMsg(clientId, stream)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		c.broadcastChan <- clientBroadcast{ClientId: clientId, resp: Chat.BroadcastResponse{Message: fmt.Sprintf("%s : %s", clientId, req.Message)}}
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

func (c *ChatServer) sendBroadcastMsg(clientId string, stream Chat.Chat_BroadcastServer) {
	for {
		c.clientStreamsLock.RLock()
		cc, ok := c.clients[clientId]
		c.clientStreamsLock.RUnlock()
		if !ok {
			log.Printf("Client not registered %s", clientId)
			return
		}
		select {
		case <-stream.Context().Done():
			return
		case msg := <-cc:
			stream.Send(&msg)
		}
	}
}

func (c *ChatServer) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	Chat.RegisterChatServer(srv, c)
	srv.Serve(lis)
}
