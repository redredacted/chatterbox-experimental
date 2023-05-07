package main

import (
	ctx "context"
	"net"

	pb "github.com/RemoteENv-Team/go.grpc/protos"
	"google.golang.org/grpc"
)

type server struct {
	pb.LoginServer
}

func (s *server) Login(ctx ctx.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Message: "not implemented", Success: false, Token: "token"}, nil 
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}