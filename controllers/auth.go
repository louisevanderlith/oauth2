package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxo"
	"github.com/ory/fosite"
	"log"
	"net/http"
)

func Auth(c *gin.Context) {
	// This context will be passed to all methods.
	ctx := fosite.NewContext()

	// Let's create an AuthorizeRequest object!
	// It will analyze the request and extract important information like scopes, response type and others.
	ar, err := authProvider.NewAuthorizeRequest(ctx, c.Request)
	if err != nil {
		log.Printf("Error occurred in NewAuthorizeRequest: %+v", err)
		authProvider.WriteAuthorizeError(c.Writer, ar, err)
		return
	}

	// You have now access to authorizeRequest, Code ResponseTypes, Scopes ...

	sess := ar.GetSession()

	if sess == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	//Consent
	c.Request.ParseForm()
	if len(c.Request.PostForm["scopes"]) == 0 {
		c.HTML(http.StatusOK, "auth.html", droxo.Wrap("Auth", ar.GetRequestedScopes()))
		return
	}

	// let's see what scopes the user gave consent to
	for _, scope := range c.Request.PostForm["scopes"] {
		ar.GrantScope(scope)
	}

	// Now that the user is authorized, we set up a session:
	mySessionData := newSession(sess.GetSubject(), ar.GetGrantedAudience())

	// When using the HMACSHA strategy you must use something that implements the HMACSessionContainer.
	// It brings you the power of overriding the default values.
	//
	// mySessionData.HMACSession = &strategy.HMACSession{
	//	AccessTokenExpiry: time.Now().Add(time.Day),
	//	AuthorizeCodeExpiry: time.Now().Add(time.Day),
	// }
	//

	// If you're using the JWT strategy, there's currently no distinction between access token and authorize code claims.
	// Therefore, you both access token and authorize code will have the same "exp" claim. If this is something you
	// need let us know on github.
	//
	// mySessionData.JWTClaims.ExpiresAt = time.Now().Add(time.Day)

	// It's also wise to check the requested scopes, e.g.:
	// if authorizeRequest.GetScopes().Has("admin") {
	//     http.Error(rw, "you're not allowed to do that", http.StatusForbidden)
	//     return
	// }

	// Now we need to get a response. This is the place where the AuthorizeEndpointHandlers kick in and start processing the request.
	// NewAuthorizeResponse is capable of running multiple response type handlers which in turn enables this library
	// to support open id connect.
	response, err := authProvider.NewAuthorizeResponse(ctx, ar, mySessionData)

	// Catch any errors, e.g.:
	// * unknown client
	// * invalid redirect
	// * ...
	if err != nil {
		log.Printf("Error occurred in NewAuthorizeResponse: %+v", err)
		authProvider.WriteAuthorizeError(c.Writer, ar, err)
		return
	}

	// Last but not least, send the response!
	authProvider.WriteAuthorizeResponse(c.Writer, ar, response)
}
