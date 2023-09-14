package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/katerji/gopush/db"
	"github.com/katerji/gopush/gapi"
	"github.com/katerji/gopush/handler"
	"github.com/katerji/gopush/middleware"
	gopush "github.com/katerji/gopush/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	initDB()
	go initGRPCServer()
	initWebServer()
}

func initDB() {
	client := db.GetDbInstance()
	err := client.Ping()
	if err != nil {
		panic(err)
	}
}

func initWebServer() {
	router := gin.Default()
	api := router.Group("/api")

	api.GET(handler.LandingPath, handler.LandingController)

	auth := api.Group("/auth")
	auth.POST(handler.RegisterPath, handler.RegisterHandler)
	auth.POST(handler.LoginPath, handler.LoginHandler)
	auth.POST(handler.RefreshTokenPath, handler.RefreshTokenHandler)

	api.Use(middleware.GetAuthMiddleware())

	api.GET(handler.UserInfoPath, handler.UserInfoHandler)

	err := router.Run(":85")
	if err != nil {
		panic(err)
	}
}

func initGRPCServer() {
	grpcServer := grpc.NewServer()
	s := gapi.NewServer()
	gopush.RegisterPusherServer(grpcServer, s)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}

}
