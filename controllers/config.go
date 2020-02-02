package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type openConfig struct {
	Issuer                       string   `json:"issuer"`
	AuthorizationEndpoint        string   `json:"authorization_endpoint"`
	TokenEndpoint                string   `json:"token_endpoint"`
	UserInfoEndpoint             string   `json:"userinfo_endpoint"`
	JwkURI                       string   `json:"jwks_uri"`
	ScopesSupported              []string `json:"scopes_supported"`
	ResponseTypesSupported       []string `json:"response_types_supported"`
	GrantTypesSupported          []string `json:"grant_types_supported"`
	SubjectTypesSupported        []string `json:"subject_types_supported"`
	SigningAlgsSupported         []string `json:"id_token_signing_alg_values_supported"`
	EncryptionAlgsSupported      []string `json:"id_token_encryption_alg_values_supported"`
	EncryptionEncodingsSupported []string `json:"id_token_encryption_enc_values_supported"`
	AuthMethodsSupported         []string `json:"token_endpoint_auth_methods_supported"`
	AuthSigningAlgsSupported     []string `json:"token_endpoint_auth_signing_alg_values_supported"`
	ClaimsParamSupported         bool     `json:"claims_parameter_supported"`
	RequestParamSupported        bool     `json:"request_parameter_supported"`
	RequestUriParamSupported     bool     `json:"request_uri_parameter_supported"`
}

//.well-known/openid-configuration
func OpenIDConfig(c *gin.Context) {
	issuer := fmt.Sprintf("https://oauth2.%s", _host)
	var responseTypes []string
	for _, v := range _server.Config.AllowedResponseTypes {
		responseTypes = append(responseTypes, v.String())
	}

	var grantTypes []string
	for _, v := range _server.Config.AllowedGrantTypes {
		grantTypes = append(grantTypes, v.String())
	}

	result := openConfig{
		Issuer:                 issuer,
		AuthorizationEndpoint:  fmt.Sprintf("%s/authorize", issuer),
		TokenEndpoint:          fmt.Sprintf("%s/token", issuer),
		UserInfoEndpoint:       fmt.Sprintf("%s/info", issuer),
		JwkURI:                 fmt.Sprintf("%s/jwks", issuer),
		ScopesSupported:        _scopes,
		ResponseTypesSupported: responseTypes,
		GrantTypesSupported:    grantTypes,
		SubjectTypesSupported:  []string{"public"},
		SigningAlgsSupported:   []string{"HS512"},
		EncryptionAlgsSupported: []string{
			"RSA1_5",
			"RSA-OAEP",
			"RSA-OAEP-256",
			"dir",
		},
		EncryptionEncodingsSupported: []string{"A256CBC-HS512"},
		AuthMethodsSupported:         []string{"client_secret_post"},
		AuthSigningAlgsSupported: []string{
			"HS256",
			"RS256",
		},
		ClaimsParamSupported:     false,
		RequestParamSupported:    false,
		RequestUriParamSupported: false,
	}
	c.JSON(http.StatusOK, result)
}
