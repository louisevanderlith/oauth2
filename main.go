package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/oauth2/controllers"
	"github.com/louisevanderlith/oauth2/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	certPath := "/signcerts/"
	controllers.InitOAuthServer(certPath)

	//Download latest Theme
	err := droxo.UpdateTheme(droxo.UriTheme)

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	tmpl, err := droxo.LoadTemplates("./views")

	if err != nil {
		panic(err)
	}

	r.HTMLRender = tmpl

	r.GET("/authorize", controllers.Authorize)
	r.POST("/authorize", controllers.Authorize)
	r.GET("/login", controllers.Login)
	r.POST("/login", controllers.LoginPost)
	r.GET("/consent", controllers.Consent)
	r.POST("/token", controllers.Token)
	r.GET("/jwks", controllers.GetJWKs)

	clntGrp := r.Group("/")
	accts, err := core.GetClientAccounts()

	if err != nil {
		panic(err)
	}
	clntGrp.Use(gin.BasicAuth(accts))

	clntGrp.GET("profiles", controllers.Profiles)
	clntGrp.POST("info", controllers.Info)

	err = r.Run(":8086")

	if err != nil {
		panic(err)
	}
}
