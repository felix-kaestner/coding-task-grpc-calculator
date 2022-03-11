package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/felix-kaestner/calculator/internal/app/server"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server listening at %v", lis.Addr())
	if err := server.New().Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
