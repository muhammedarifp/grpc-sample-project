package main

import (
	"fmt"
	"log"
	"net"

	pb "grpcsamp/proto"

	"context"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcServ := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServ, &helloServer{})
	fmt.Printf("starting on :: [ %s : %s ] : \n", "Port"+" : "+"%s"+port[1:])

	if grpcErr := grpcServ.Serve(lis); grpcErr != nil {
		log.Fatal(err.Error())
	}

}

func (helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
