package core

import (
	"context"
	"errors"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
	"github.com/ory/fosite"
	"log"
)

type store struct {
	Clients                husk.Tabler
	AuthorizeCodes         husk.Tabler
	IDSessions             husk.Tabler
	AccessTokens           husk.Tabler
	Implicit               husk.Tabler
	RefreshTokens          husk.Tabler
	PKCES                  husk.Tabler
	Users                  husk.Tabler
	Forgotten              husk.Tabler
	AccessTokenRequestIDs  map[string]string
	RefreshTokenRequestIDs map[string]string
}

var authStore store

func CreateContext() *store{
	authStore = store{
		Clients:                husk.NewTable(Client{}, serials.GobSerial{}),
		AuthorizeCodes:         husk.NewTable(AuthCode{}, serials.GobSerial{}),
		IDSessions:             husk.NewTable(Fosi{}, serials.GobSerial{}),
		AccessTokens:           husk.NewTable(Fosi{}, serials.GobSerial{}),
		Implicit:               husk.NewTable(Fosi{}, serials.GobSerial{}),
		RefreshTokens:          husk.NewTable(Fosi{}, serials.GobSerial{}),
		PKCES:                  husk.NewTable(Fosi{}, serials.GobSerial{}),
		Users:                  husk.NewTable(User{}, serials.GobSerial{}),
		Forgotten:              husk.NewTable(Forgot{}, serials.GobSerial{}),
		AccessTokenRequestIDs:  make(map[string]string),
		RefreshTokenRequestIDs: make(map[string]string),
	}

	seed()

	return &authStore
}

func Shutdown() {
	authStore.Clients.Save()
	authStore.AuthorizeCodes.Save()
	authStore.IDSessions.Save()
	authStore.AccessTokens.Save()
	authStore.Implicit.Save()
	authStore.RefreshTokens.Save()
	authStore.PKCES.Save()
	authStore.Users.Save()
	authStore.Forgotten.Save()
}

func seed() {
	err := authStore.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	authStore.Users.Save()

	err = authStore.Clients.Seed("db/clients.seed.json")

	if err != nil {
		panic(err)
	}

	authStore.Clients.Save()
}

func (s store) CreateOpenIDConnectSession(_ context.Context, authorizeCode string, requester fosite.Requester) error {
	obj := Fosi{authorizeCode, requester}
	cset := s.IDSessions.Create(obj)

	if cset.Error != nil {
		return cset.Error
	}

	return s.IDSessions.Save()
}

func (s *store) GetOpenIDConnectSession(_ context.Context, authorizeCode string, requester fosite.Requester) (fosite.Requester, error) {
	rec, err := s.IDSessions.FindFirst(byID(authorizeCode))

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	f := rec.Data().(Fosi)

	return f.Requester, nil
}

func (s *store) DeleteOpenIDConnectSession(_ context.Context, authorizeCode string) error {
	rec, err := s.IDSessions.FindFirst(byID(authorizeCode))

	if err != nil {
		return err
	}

	err = s.IDSessions.Delete(rec.GetKey())

	if err != nil {
		return err
	}

	return s.IDSessions.Save()
}

func (s *store) GetClient(_ context.Context, id string) (fosite.Client, error) {
	rec, err := s.Clients.FindFirst(byClientID(id))

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	return rec.Data().(fosite.Client), nil
}

func (s *store) CreateAuthorizeCodeSession(_ context.Context, code string, req fosite.Requester) error {
	cset := s.AuthorizeCodes.Create(AuthCode{true, code, req})

	if cset.Error != nil {
		return cset.Error
	}

	return 	s.AuthorizeCodes.Save()
}

func (s *store) GetAuthorizeCodeSession(_ context.Context, code string, _ fosite.Session) (fosite.Requester, error) {
	rec, err := s.AuthorizeCodes.FindFirst(byAuthCode(code))

	if err != nil {
		log.Println(err)
		return nil, fosite.ErrInvalidatedAuthorizeCode
	}

	ac := rec.Data().(AuthCode)
	return ac.Requester, nil
}

func (s *store) InvalidateAuthorizeCodeSession(ctx context.Context, code string) error {
	rec, err := s.AuthorizeCodes.FindFirst(byAuthCode(code))

	if err != nil {
		log.Println(err)
		return fosite.ErrNotFound
	}

	data := rec.Data().(AuthCode)
	data.Active = false
	err = rec.Set(data)

	if err != nil {
		return err
	}

	err = s.AuthorizeCodes.Update(rec)

	if err != nil {
		return err
	}

	return s.AuthorizeCodes.Save()
}

func (s *store) DeleteAuthorizeCodeSession(_ context.Context, code string) error {
	rec, err := s.AuthorizeCodes.FindFirst(byAuthCode(code))

	if err != nil {
		log.Println(err)
		return fosite.ErrNotFound
	}

	err = s.AuthorizeCodes.Delete(rec.GetKey())

	if err != nil {
		return err
	}

	return s.AuthorizeCodes.Save()
}

func (s *store) CreatePKCERequestSession(_ context.Context, code string, req fosite.Requester) error {
	obj := Fosi{code, req}
	cset := s.PKCES.Create(obj)

	if cset.Error != nil {
		return cset.Error
	}

	return s.PKCES.Save()
}

func (s *store) GetPKCERequestSession(_ context.Context, code string, _ fosite.Session) (fosite.Requester, error) {
	rec, err := s.PKCES.FindFirst(byID(code))

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	f := rec.Data().(Fosi)

	return f.Requester, nil
}

func (s *store) DeletePKCERequestSession(_ context.Context, code string) error {
	rec, err := s.PKCES.FindFirst(byID(code))

	if err != nil {
		return err
	}

	err = s.PKCES.Delete(rec.GetKey())

	if err != nil {
		return err
	}

	return s.PKCES.Save()
}

func (s *store) CreateAccessTokenSession(_ context.Context, signature string, req fosite.Requester) error {
	obj := Fosi{signature, req}
	cset := s.AccessTokens.Create(obj)

	if cset.Error != nil {
		return cset.Error
	}

	err := s.AccessTokens.Save()

	if err != nil {
		return nil
	}

	s.AccessTokenRequestIDs[req.GetID()] = signature

	return nil
}

func (s *store) GetAccessTokenSession(_ context.Context, signature string, _ fosite.Session) (fosite.Requester, error) {
	rec, err := s.AccessTokens.FindFirst(byID(signature))

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	f := rec.Data().(Fosi)

	return f.Requester, nil
}

func (s *store) DeleteAccessTokenSession(_ context.Context, signature string) error {
	rec, err := s.AccessTokens.FindFirst(byID(signature))

	if err != nil {
		return err
	}

	err = s.AccessTokens.Delete(rec.GetKey())

	if err != nil {
		return err
	}

	return s.AccessTokens.Save()
}

func (s *store) CreateRefreshTokenSession(_ context.Context, signature string, req fosite.Requester) error {
	obj := Fosi{signature, req}
	cset := s.RefreshTokens.Create(obj)

	if cset.Error != nil {
		return cset.Error
	}

	err := s.RefreshTokens.Save()

	if err != nil {
		return nil
	}

	s.RefreshTokenRequestIDs[req.GetID()] = signature

	return nil
}

func (s *store) GetRefreshTokenSession(_ context.Context, signature string, _ fosite.Session) (fosite.Requester, error) {
	rec, err := s.RefreshTokens.FindFirst(byID(signature))

	if err != nil {
		return nil, fosite.ErrNotFound
	}

	f := rec.Data().(Fosi)

	return f.Requester, nil
}

func (s *store) DeleteRefreshTokenSession(_ context.Context, signature string) error {
	rec, err := s.RefreshTokens.FindFirst(byID(signature))

	if err != nil {
		return err
	}

	err = s.RefreshTokens.Delete(rec.GetKey())

	if err != nil {
		return err
	}

	return s.RefreshTokens.Save()
}

func (s *store) CreateImplicitAccessTokenSession(_ context.Context, code string, req fosite.Requester) error {
	obj := Fosi{code, req}
	cset := s.Implicit.Create(obj)

	if cset.Error != nil {
		return cset.Error
	}

	return s.Implicit.Save()
}

func (s *store) Authenticate(_ context.Context, name string, secret string) error {
		_, err := Login(name, secret)

		if err != nil {
			log.Println(err)
			return errors.New("Invalid credentials")
		}

		return  nil
}

func (s *store) RevokeRefreshToken(ctx context.Context, requestID string) error {
	if signature, exists := s.RefreshTokenRequestIDs[requestID]; exists {
		s.DeleteRefreshTokenSession(ctx, signature)
		s.DeleteAccessTokenSession(ctx, signature)
	}

	return nil
}

func (s *store) RevokeAccessToken(ctx context.Context, requestID string) error {
	if signature, exists := s.AccessTokenRequestIDs[requestID]; exists {
		s.DeleteAccessTokenSession(ctx, signature)
	}

	return nil
}
