package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/oauth2/controllers"
	"github.com/louisevanderlith/oauth2/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	certPath := "/signcerts/"
	controllers.InitOAuthServer(certPath)
	r := gin.Default()

	r.GET("/auth", controllers.Auth)
	r.GET("/authorize", controllers.Authorize)
	r.GET("/login", controllers.Login)
	r.POST("/login", controllers.LoginPost)
	r.POST("/token", controllers.Token)
	r.GET("/clients", controllers.Clients)
	r.POST("/info", controllers.Info	)

	err := r.Run(":8086")

	if err != nil {
		panic(err)
	}
}
