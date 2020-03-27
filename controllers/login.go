package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/oauth2/core"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		log.Println("no session store", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); ok {
		c.Header("Location", "/consent")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}

	c.HTML(http.StatusOK, "login.html", droxo.Wrap("Login", nil))
}

func LoginPost(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		log.Println("no session store", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	r := c.Request

	if r.Form == nil {
		if err := r.ParseForm(); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	k, err := core.Login(r.Form.Get("username"), r.Form.Get("password"))

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	store.Set("LoggedInUserID", k)
	store.Save()

	c.Header("Location", "/consent")
	c.Writer.WriteHeader(http.StatusFound)
}
