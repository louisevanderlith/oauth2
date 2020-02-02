package controllers

import (
	"crypto/x509"
	"github.com/louisevanderlith/oauth2/core"
	"github.com/louisevanderlith/oauth2/signing"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-session/session"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var (
	_server *server.Server
	_host string
	_scopes []string
)

func InitOAuthServer(certPath string, host string) {
	_host = host
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.SetClientTokenCfg(manage.DefaultClientTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	err := signing.Initialize(certPath)

	if err != nil {
		panic(err)
	}

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate(x509.MarshalPKCS1PrivateKey(signing.PrivateKey), jwt.SigningMethodHS512))
	manager.MapClientStorage(core.NewClientStore())

	_server = server.NewServer(server.NewConfig(), manager)
	_server.SetAllowGetAccessRequest(true)
	_server.SetClientInfoHandler(server.ClientFormHandler)
	_server.SetPasswordAuthorizationHandler(core.Login)
	_server.SetUserAuthorizationHandler(userAuthorizeHandler)

	_server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	_server.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	_scopes = []string{
		"openid",
		"offline_access",
		"profile",
		"artifact",
		"comms",
		"comment",
		"blog",
		"theme",
		"vin",
	}
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		return "", err
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return "", nil
	}

	userID := uid.(string)
	store.Delete("LoggedInUserID")
	err = store.Save()

	if err != nil {
		return "", err
	}

	return userID, nil
}
