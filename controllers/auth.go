package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/louisevanderlith/droxo"
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
	"time"
)

func Auth_old(c *gin.Context) {
	store, err := session.Start(nil, c.Writer, c.Request)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "auth.html", droxo.Wrap("Auth", nil))
}

func Auth(c *gin.Context){
	ctx := fosite.NewContext()

	authorizeRequest, err := authProvider.NewAuthorizeRequest(ctx, c.Request)
	if err != nil {
		log.Println(err)
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	session =  newSession(authorizeRequest.)
	if errors.Cause(err) == consent.ErrAbortOAuth2Request {
		// do nothing
		return
	} else if err != nil {
		x.LogError(err, h.r.Logger())
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	for _, scope := range session.GrantedScope {
		authorizeRequest.GrantScope(scope)
	}

	for _, audience := range session.GrantedAudience {
		authorizeRequest.GrantAudience(audience)
	}

	openIDKeyID, err := h.r.OpenIDJWTStrategy().GetPublicKeyID(r.Context())
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	var accessTokenKeyID string
	if h.c.AccessTokenStrategy() == "jwt" {
		accessTokenKeyID, err = h.r.AccessTokenJWTStrategy().GetPublicKeyID(r.Context())
		if err != nil {
			x.LogError(err, h.r.Logger())
			h.writeAuthorizeError(w, r, authorizeRequest, err)
			return
		}
	}

	authorizeRequest.SetID(session.Challenge)

	claims := &jwt.IDTokenClaims{
		Subject:                             session.ConsentRequest.SubjectIdentifier,
		Issuer:                              strings.TrimRight(h.c.IssuerURL().String(), "/") + "/",
		IssuedAt:                            time.Now().UTC(),
		AuthTime:                            session.AuthenticatedAt,
		RequestedAt:                         session.RequestedAt,
		Extra:                               session.Session.IDToken,
		AuthenticationContextClassReference: session.ConsentRequest.ACR,

		// We do not need to pass the audience because it's included directly by ORY Fosite
		// Audience:    []string{authorizeRequest.GetClient().GetID()},

		// This is set by the fosite strategy
		// ExpiresAt:   time.Now().Add(h.IDTokenLifespan).UTC(),
	}
	claims.Add("sid", session.ConsentRequest.LoginSessionID)

	// done
	response, err := h.r.OAuth2Provider().NewAuthorizeResponse(ctx, authorizeRequest, &Session{
		DefaultSession: &openid.DefaultSession{
			Claims: claims,
			Headers: &jwt.Headers{Extra: map[string]interface{}{
				// required for lookup on jwk endpoint
				"kid": openIDKeyID,
			}},
			Subject: session.ConsentRequest.Subject,
		},
		Extra:            session.Session.AccessToken,
		KID:              accessTokenKeyID,
		ClientID:         authorizeRequest.GetClient().GetID(),
		ConsentChallenge: session.Challenge,
	})
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	h.r.OAuth2Provider().WriteAuthorizeResponse(w, authorizeRequest, response)
}