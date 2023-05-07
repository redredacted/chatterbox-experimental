package service

import (
	ctx "context"

	pb "github.com/RemoteENv-Team/go.grpc/protos"
	"github.com/golang-jwt/jwt/v5"
)

type LoginService struct {
	pb.LoginServer
}

func (s *LoginService) Login(ctx ctx.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
	})

	tokenString, err := token.SignedString([]byte("secret")); if err != nil {
		return &pb.LoginResponse{Message: "failed to generate token", Success: false, Token: ""}, err
	}

	return &pb.LoginResponse{Message: "you got a bunk token congrats!", Success: true, Token: tokenString}, nil
}
