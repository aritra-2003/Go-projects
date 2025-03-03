package main

import (
	"context"
	"log"
	"time"

	 pb "example.com/crud/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a new user
	res, err := client.CreateUser(ctx, &pb.CreateUserRequest{Name: "John Doe", Email: "john@example.com"})
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	log.Printf("User created: %v", res)
}