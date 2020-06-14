package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/at8109/go-grpc/chat"
)

func main(){
	lis , err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}


	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer , &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve gRPC server over Port 9000: %v", err)
	}
}