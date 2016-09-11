package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"context"
	pb "github.com/Arthelon/grpc-fake/proto"
)

type fakeGeneratorServer struct {

}

func (fg *fakeGeneratorServer) GetUser(ctx context.Context, _ *pb.EmptyMessage) {

}

func (fg *fakeGeneratorServer) GetAddress(ctx context.Context, _ *pb.EmptyMessage) {

}

func (fg *fakeGeneratorServer) GetEmail(ctx context.Context, _ *pb.EmptyMessage) {

}

func (fg *fakeGeneratorServer) GetDate(ctx context.Context, _ *pb.EmptyMessage) {

}


func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		grpclog.Fatalf("Failed to listen to port 3000: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFakeGeneratorServer(&grpcServer, fakeGeneratorServer{})
	grpcServer.Serve(listener)
}