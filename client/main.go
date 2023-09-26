package main

import (
	"context"
	pb "grpcsamp/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	callSayHello(client)
}

func callSayHello(client pb.GreetServiceClient) {
	ctx, close := context.WithTimeout(context.Background(), time.Second)
	defer close()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatal("Resp Err -- ", err.Error())
	}

	log.Print("%v", res.Message)
}
