package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"golang.org/x/net/context"
	pb "github.com/Arthelon/grpc-fake/proto"
	"github.com/icrowley/fake"
)

type fakeGeneratorServer struct {}

func generateAddress() pb.Address {
	return pb.Address{
		Continent: fake.Continent(),
		Country: fake.Country(),
		City: fake.City(),
		State: fake.State(),
		StreetAddress: fake.StreetAddress(),
		Zip: fake.Zip(),
		Latitude: fake.Latitute(),
		Longitude: fake.Longitude(),
	}
}

func generateDate() pb.Date {
	return pb.Date{
		Year: int64(fake.Year(1000, 2016)),
		Month: int64(fake.MonthNum()),
		Day: int64(fake.Day()),
	}
}

func (fg *fakeGeneratorServer) GetUser(ctx context.Context, in *pb.EmptyMessage) (*pb.User, error) {
	date := generateDate()
	address := generateAddress()
	return &pb.User{
		FullName: fake.FullName(),
		Username: fake.UserName(),
		Password: fake.SimplePassword(),
		Email: fake.EmailAddress(),
		PhoneNumber: fake.Phone(),
		Gender: fake.Gender(),
		Language: fake.Language(),
		Birthday: &date,
		Address: &address,
	}, nil
}

func (fg *fakeGeneratorServer) GetAddress(ctx context.Context, in *pb.EmptyMessage) (*pb.Address, error) {
	address := generateAddress()
	return &address, nil
}

func (fg *fakeGeneratorServer) GetEmail(ctx context.Context, in *pb.EmptyMessage) (*pb.Email, error) {
	return &pb.Email{
		Title: fake.EmailSubject(),
		Body: fake.EmailBody(),
		From: fake.EmailAddress(),
		To: fake.EmailAddress(),
	}, nil
}

func (fg *fakeGeneratorServer) GetDate(ctx context.Context, in *pb.EmptyMessage) (*pb.Date, error) {
	date := generateDate()
	return &date, nil
}


func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		grpclog.Fatalf("Failed to listen to port 3000: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFakeGeneratorServer(grpcServer, &fakeGeneratorServer{})
	grpcServer.Serve(listener)
}