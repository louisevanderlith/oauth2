package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"net/http"
	"net/url"
)

func Authorize(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}

	c.Request.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = _server.HandleAuthorizeRequest(c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, nil)
}
