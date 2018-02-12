package main

import (
	"fmt"
	"github.com/gitsen/playground/chat/server"
)

func main() {
	c := server.New()
	fmt.Println("Starting Chat Server")
	c.Run()
}
