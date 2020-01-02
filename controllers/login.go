package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPost(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	r := c.Request

	if r.Form == nil {
		if err := r.ParseForm(); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}

	store.Set("LoggedInUserID", r.Form.Get("username"))
	store.Save()

	c.Redirect(http.StatusFound, "/auth")
}
