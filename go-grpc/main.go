package main

import (
	"net"

	pb "github.com/RemoteENv-Team/go.grpc/protos"
	"github.com/RemoteENv-Team/go.grpc/service"
	"github.com/RemoteENv-Team/go.grpc/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	netListen, err := net.Listen("tcp", ":50051"); if err != nil {
		tools.LogStdout().Fatal("Failed to listen: " + err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterLoginServer(s, &service.LoginService{})
	
	// Register reflection service on gRPC server.
	reflection.Register(s); if err := s.Serve(netListen); err != nil {
		tools.LogStdout().Fatal("Failed to serve: " + err.Error())
	}

	tools.LogStdout().Info("Successfully started gRPC server!")
}
