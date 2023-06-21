package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const somethingWentWrongMessage = "Something went wrong."

func sendBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Bad Request.",
	})
}

func sendBadRequestWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

func sendError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": somethingWentWrongMessage,
	})
}

func sendErrorMessage(c *gin.Context, message string) {
	if message == "" {
		message = somethingWentWrongMessage
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
	})
}

func sendResponseMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func sendJSONResponse(c *gin.Context, json any) {
	c.JSON(http.StatusOK, json)
}
