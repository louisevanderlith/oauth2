package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Introspection endpoint
func Info(c *gin.Context) {
	accCode, ok := c.GetPostForm("access_code")

	if !ok {
		log.Println("no access_code")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	info, err := _server.Manager.LoadAccessToken(accCode)

	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, info)
}
