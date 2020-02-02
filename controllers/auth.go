package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/louisevanderlith/droxo"
	"net/http"
)

func Auth(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "auth.html", droxo.Wrap("Auth", nil))
}
