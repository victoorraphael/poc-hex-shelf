package main

import (
	"context"
	pb "github.com/victoorraphael/poc-hex-shelf/internal/handlers/hellogrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

var addr = "0.0.0.0:7070"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)
	callHello(c)
}

func callHello(client pb.HelloServiceClient) {
	res, err := client.Hello(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetMsg())
}
