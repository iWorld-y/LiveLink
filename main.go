package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/iWorld-y/TradeHub/idl/pb/link"
	"github.com/iWorld-y/TradeHub/internal/handler"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	server := grpc.NewServer()
	link.RegisterLinkServer(server, &handler.LinkHandler{})
	log.Println("server start at 127.0.0.1:8080")

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Println(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	gwmux := runtime.NewServeMux()
	if err := link.RegisterLinkHandler(context.Background(), gwmux, conn); err != nil {
		log.Println(err)
		return
	}
	gwServer := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: gwmux,
	}
	log.Println("gateway start at 127.0.0.1:8081")
	if err := gwServer.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
