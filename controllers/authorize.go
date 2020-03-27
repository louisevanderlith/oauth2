package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"log"
	"net/http"
	"net/url"
)

func Authorize(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		log.Println("no session store", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Abort()
}
