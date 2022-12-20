package hellogrpc

import (
	"context"
	"github.com/victoorraphael/poc-hex-shelf/internal/core/ports"
	pb "github.com/victoorraphael/poc-hex-shelf/internal/handlers/hellogrpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	srv ports.HelloService
	pb.HelloServiceServer
}

func NewGRPCServer(service ports.HelloService) *server {
	return &server{srv: service}
}

func (s *server) Hello(_ context.Context, _ *emptypb.Empty) (*pb.HelloResponse, error) {
	msg := s.srv.Get()
	return &pb.HelloResponse{Msg: msg}, nil
}
