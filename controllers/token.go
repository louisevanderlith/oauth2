package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Token(c *gin.Context) {
	err := _server.HandleTokenRequest(c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, nil)
}
