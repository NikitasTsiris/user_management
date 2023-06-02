package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	pb "user_management/usermgmt"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, newUser *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", newUser.GetName())

	var user_id int32 = int32(rand.Intn(1000))

	return &pb.User{Name: newUser.GetName(), Age: newUser.GetAge(), Id: user_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserManagementServer(server, &UserManagementServer{})

	log.Printf("Server Listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
