package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "github.com/Arthelon/grpc-fake/proto"
	"fmt"
	"golang.org/x/net/context"
)


func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("Error occured while dialing server: %v", err)
		return
	}
	fmt.Println("Connected to server")
	defer conn.Close()
	client := pb.NewFakeGeneratorClient(conn)
	addr, err := client.GetAddress(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		grpclog.Fatalf("Error while running RPC: %v\n", err)
		return
	}
	fmt.Printf("Address: %v\n", addr)
	date, err := client.GetDate(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		grpclog.Fatalf("Error while running RPC: %v\n", err)
		return
	}
	fmt.Printf("Date: %v\n", date)
	email, err := client.GetEmail(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		grpclog.Fatalf("Error while running RPC: %v\n", err)
		return
	}
	fmt.Printf("Email: %v\n", email)
	user, err := client.GetUser(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		grpclog.Fatalf("Error while running RPC: %v\n", err)
		return
	}
	fmt.Printf("User: %v\n", user)
	fmt.Println("Finished Executing")
}