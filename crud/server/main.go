package main

import (
	"context"
	"log"
	"net"

	pb "example.com/crud/proto"
	"example.com/crud/db"

	"google.golang.org/grpc"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

// Create User
func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := db.User{Name: req.Name, Email: req.Email}
	db.DB.Create(&user)

	return &pb.UserResponse{Id: int32(user.ID), Name: user.Name, Email: user.Email}, nil
}

// Get User by ID
func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	var user db.User
	result := db.DB.First(&user, req.Id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.UserResponse{Id: int32(user.ID), Name: user.Name, Email: user.Email}, nil
}

// Update User
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	var user db.User
	db.DB.First(&user, req.Id)

	user.Name = req.Name
	user.Email = req.Email
	db.DB.Save(&user)

	return &pb.UserResponse{Id: int32(user.ID), Name: user.Name, Email: user.Email}, nil
}

// Delete User
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	db.DB.Delete(&db.User{}, req.Id)
	return &pb.DeleteUserResponse{Message: "User deleted successfully"}, nil
}

func main() {
	db.InitDB()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{})

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}