package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "user_management/usermgmt"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	var result = "Y"
	var name string
	var age int32

	for result == "Y" || result == "y" {
		fmt.Println("resutl: ", result)

		fmt.Print("User's Name? ... ")
		fmt.Scan(&name)

		fmt.Print("User's Age? ... ")
		fmt.Scan(&age)

		response, err := client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})

		if err != nil {
			log.Fatalf("Could not create user: %v", err)
		}

		log.Printf(`Successfully Added User With Details:
		NAME: %s
		AGE: %d
		ID: %d`, response.GetName(), response.GetAge(), response.GetId())

		fmt.Printf("Enter user? Y/N ... ")
		fmt.Scan(&result)

		fmt.Println("resutl: ", result)
	}
}
