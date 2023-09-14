package gapi

import (
	"context"
	"github.com/katerji/gopush/input"
	gopush "github.com/katerji/gopush/proto"
	"github.com/katerji/gopush/service"
)

type Server struct {
	gopush.UnimplementedPusherServer
}

func (s Server) Login(_ context.Context, request *gopush.LoginRequest) (*gopush.LoginResponse, error) {
	authService := service.AuthService{}
	authInput := input.AuthInput{
		Email:    request.Email,
		Password: request.Password,
	}
	user, err := authService.Login(authInput)
	return &gopush.LoginResponse{
		User:         user,
		Token:        "",
		RefreshToken: "",
	}, err
}

func (s Server) Register(_ context.Context, request *gopush.RegisterRequest) (*gopush.GenericResponse, error) {
	registerUserInput := input.AuthInput{
		Email:    request.Email,
		Password: request.Password,
	}
	userService := service.AuthService{}
	_, err := userService.Register(registerUserInput)
	if err != nil {
		return &gopush.GenericResponse{
			Success: false,
			Message: "Email already exists",
		}, nil
	}
	return &gopush.GenericResponse{
		Success: true,
	}, nil
}

func NewServer() *Server {
	return &Server{}
}
