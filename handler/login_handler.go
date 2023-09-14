package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/gopush/input"
	gopush "github.com/katerji/gopush/proto"
	"github.com/katerji/gopush/service"
)

const LoginPath = "/login"

func LoginHandler(c *gin.Context) {
	request := &gopush.LoginRequest{}
	err := c.BindJSON(request)
	if err != nil {
		sendBadRequest(c)
		return
	}
	authService := service.AuthService{}
	authInput := input.AuthInput{
		Email:    request.Email,
		Password: request.Password,
	}
	user, err := authService.Login(authInput)
	if err != nil {
		sendErrorMessage(c, err.Error())
		return
	}
	jwtService := service.JWTService{}
	token, err := jwtService.CreateJwt(user)
	if err != nil {
		sendErrorMessage(c, "")
		return
	}
	refreshToken, err := jwtService.CreateRefreshJwt(user)
	if err != nil {
		sendErrorMessage(c, "")
		return
	}
	response := &gopush.LoginResponse{
		User:         user,
		Token:        token,
		RefreshToken: refreshToken,
	}
	sendJSONResponse(c, response)
	return
}
