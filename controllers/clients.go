package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/oauth2/core"
	"net/http"
)

func Clients(c *gin.Context) {
	result, err := core.GetAllClients()

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
