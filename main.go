package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/oauth2/controllers"
	"github.com/louisevanderlith/oauth2/core"
	"os"
)

func main() {
	host := os.Getenv("HOST")

	if len(host) == 0 {
		panic(errors.New("no host provided"))
		return
	}

	prof := os.Getenv("PROFILE")

	if len(prof) == 0 {
		panic(errors.New("no profile provided"))
		return
	}

	core.CreateContext()
	defer core.Shutdown()

	certPath := "/signcerts/"
	controllers.InitProvider(certPath, host)

	droxo.AssignOperator(prof, host)
	//Download latest Theme
	err := droxo.UpdateTheme("http://theme:8093")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	tmpl, err := droxo.LoadTemplates("./views")

	if err != nil {
		panic(err)
	}

	r.HTMLRender = tmpl

	r.GET("/auth", controllers.Auth)
	r.POST("/authorize", controllers.Authorize)
	r.GET("/authorize", controllers.Authorize)
	r.GET("/login", controllers.Login)
	r.POST("/login", controllers.LoginPost)
	r.POST("/token", controllers.Token)
	r.GET("/clients", controllers.Clients)
	r.POST("/info", controllers.Info	)
	r.GET("/.well-known/openid-configuration", controllers.OpenIDConfig)
	r.GET("/keys", controllers.GetJWKs)
	err = r.Run(":8086")

	if err != nil {
		panic(err)
	}
}
