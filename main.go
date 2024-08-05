package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/iWorld-y/LiveLink/idl/pb/link"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println(err)
		return
	}
	server := grpc.NewServer()
	link.RegisterLinkServer(server)
}
