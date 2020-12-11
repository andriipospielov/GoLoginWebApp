package main

import (
	"github.com/andriipospielov/LoginWebApp/controller"
	"github.com/andriipospielov/LoginWebApp/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	accountController = *controller.AccountController{}.New()
	authMiddleWare    = middleware.NewJwtAuthenticator()
)

func main() {

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())
	defer accountController.AccountRepository.CloseConnection()
	errInit := authMiddleWare.MiddlewareInit()

	if errInit != nil {
		log.Fatal("JWT Error:" + errInit.Error())
	}

	server.POST("/account", accountController.Create)
	server.POST("/account/login", authMiddleWare.LoginHandler)
	server.POST("/account/logout", authMiddleWare.LogoutHandler)

	auth := server.Group("/auth")
	auth.Use(authMiddleWare.MiddlewareFunc())
	auth.GET("/refresh_token", authMiddleWare.RefreshHandler)
	auth.GET("/account", accountController.Index)
	auth.PUT("/account/:id", accountController.Update)

	err := server.Run("127.0.0.1:8080")

	if err != nil {
		log.Fatal(err)
	}
}
