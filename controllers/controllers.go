package controllers

import (
	"fmt"
	"github.com/louisevanderlith/oauth2/core"
	"github.com/louisevanderlith/oauth2/signing"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
	"gopkg.in/oauth2.v3/server"
	"time"
)

var (
	_server *server.Server
	_host   string
	_scopes []string

	authProvider fosite.OAuth2Provider
	strat compose.CommonStrategy
)

func InitProvider(certPath, host string) {
	_host= host
	err := signing.Initialize(certPath)

	if err != nil {
		panic(err)
	}

	cfg := &compose.Config{}
	store := core.CreateContext()

	strat = compose.CommonStrategy{
		CoreStrategy:               nil,
		OpenIDConnectTokenStrategy: compose.NewOpenIDConnectStrategy(cfg, signing.PrivateKey),
		JWTStrategy:                compose.NewOAuth2JWTStrategy(signing.PrivateKey, compose.NewOAuth2HMACStrategy(cfg, []byte(""), nil)),
	}
	authProvider = compose.Compose(cfg, store, strat, nil,
		compose.OAuth2AuthorizeExplicitFactory,
		compose.OAuth2AuthorizeImplicitFactory,
		compose.OAuth2ClientCredentialsGrantFactory,
		compose.OAuth2RefreshTokenGrantFactory,
		compose.OAuth2ResourceOwnerPasswordCredentialsFactory,

		compose.OAuth2TokenRevocationFactory,
		compose.OAuth2TokenIntrospectionFactory,

		// be aware that open id connect factories need to be added after oauth2 factories to work properly.
		compose.OpenIDConnectExplicitFactory,
		compose.OpenIDConnectImplicitFactory,
		compose.OpenIDConnectHybridFactory,
		compose.OpenIDConnectRefreshFactory, )
}

func newSession(userKey string, audience []string) *openid.DefaultSession {
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      fmt.Sprintf("https://oauth2.%s", _host),
			Subject:     userKey,
			Audience:    audience,
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}