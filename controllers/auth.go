package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"net/http"
)

func Auth(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		c.Redirect(http.StatusFound, "/login")
	}

	c.JSON(http.StatusOK, nil)
}
