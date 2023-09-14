package handler

import (
	"github.com/gin-gonic/gin"
	gopush "github.com/katerji/gopush/proto"
)

const UserInfoPath = "/user"

type UserInfoResponse struct {
	User *gopush.User `json:"user"`
}

func UserInfoHandler(c *gin.Context) {
	user := c.MustGet("user").(*gopush.User)
	response := &UserInfoResponse{
		User: user,
	}
	sendJSONResponse(c, response)
	return
}
