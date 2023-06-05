package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
	"context"
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

	reader := bufio.NewReader(os.Stdin)

	for result == "Y" || result == "y" {
		fmt.Print("User's Name? ... ")
		name, _= reader.ReadString('\n')
		name = strings.Replace(name, "\n", "", -1)

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

		for {

			fmt.Printf("Enter user? Y/N ... ")
			fmt.Scan(&result)

			if(strings.Contains("yYnN", result)){
				break
			}
		}
	}
}
