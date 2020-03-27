package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/louisevanderlith/droxo"
	"log"
	"net/http"
)

func Consent(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)
	if err != nil {
		log.Println("no session store", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userId, ok := store.Get("LoggedInUserID")

	if !ok {
		c.Header("Location", "/login")
		c.Writer.WriteHeader(http.StatusFound)
		return
	}

	c.HTML(http.StatusOK, "consent.html", droxo.Wrap("Consent", userId))
}
