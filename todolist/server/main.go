package main

import (
	"fmt"
	"log"
	"net"

	pb "example.com/Todo/todolist/proto"
	"example.com/Todo/ai"
	"example.com/Todo/database"

	"google.golang.org/grpc"
)

func main() {
	// Initialize Database and AI
	database.InitDB()
	ai.InitAI()

	// Start gRPC Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &TodoServiceServer{})

	fmt.Println("gRPC Server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}