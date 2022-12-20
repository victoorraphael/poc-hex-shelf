package main

import (
	"github.com/victoorraphael/poc-hex-shelf/internal/core/service/hellosrv"
	"github.com/victoorraphael/poc-hex-shelf/internal/handlers/hellogrpc"
	pb "github.com/victoorraphael/poc-hex-shelf/internal/handlers/hellogrpc/proto"
	"github.com/victoorraphael/poc-hex-shelf/internal/repositories/hellorepo"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr = "0.0.0.0:7070"

func main() {
	repo := hellorepo.NewRepo()
	service := hellosrv.New(repo)
	grpcServer := hellogrpc.NewGRPCServer(service)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, grpcServer)

	log.Println("listening on port :7070")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
