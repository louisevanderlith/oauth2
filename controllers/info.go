package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Introspection endpoint
func Info(c *gin.Context) {
	err := c.Request.ParseForm()

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	accesstoke := c.PostForm("token")

	if len(accesstoke) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	info, err := _server.Manager.LoadAccessToken(accesstoke)

	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, info)
}
