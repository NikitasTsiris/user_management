package main

import (
	"context"
	"log"
	"time"
	pb "user_management/usermgmt"

	"google.golang.org/grpc"
)

const (
	address = "user-management.default.192.168.1.240.sslip.io:80"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)

	new_users["Alice"] = 43
	new_users["Bob"] = 30

	for name, age := range new_users {
		response, err := client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})

		if err != nil {
			log.Fatalf("Could not create user: %v", err)
		}

		log.Printf(`User Details:
		NAME: %s
		AGE: %d
		ID: %d`, response.GetName(), response.GetAge(), response.GetId())
	}
}
