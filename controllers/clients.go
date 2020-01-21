package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/oauth2/core"
	"net/http"
)

func Clients(c *gin.Context) {
	result := core.GetAllClients()
	c.JSON(http.StatusOK, result)
}
