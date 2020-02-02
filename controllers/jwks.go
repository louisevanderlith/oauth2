package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/oauth2/signing"
	"math/big"
	"net/http"
)

type jwk struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Alg string `json"alg"`
	Kid string `json:"kid"`
	PublicN *big.Int `json:"n"`
	PublicE int `json:"e"`
}

func GetJWKs(c *gin.Context) {
	var keys []jwk

	pubsig := jwk{
		Kty: "RSA",
		Kid: "default",
		Use: "sig",
		Alg: "RS256",
		PublicN: signing.PrivateKey.N,
		PublicE: signing.PrivateKey.E,
	}
	keys = append(keys, pubsig)

	c.JSON(http.StatusOK, keys)
}
