package main

import (
	"fmt"
	"github.com/gitsen/playground/server"
)

func main() {
	c := server.New()
	fmt.Println("Starting Chat Server")
	c.Run()
}
