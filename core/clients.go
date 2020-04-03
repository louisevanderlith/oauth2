package core

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/husk"
	"gopkg.in/oauth2.v3"
	"strings"
)

type Clients struct{}

func NewClientStore() oauth2.ClientStore {
	return &Clients{}
}

func GetAllProfiles() (husk.Collection, error) {
	return ctx.Profiles.Find(1, 20, husk.Everything())
}

func GetClientAccounts() (gin.Accounts, error) {
	profls, err := GetAllProfiles()

	if err != nil {
		return nil, err
	}

	result := make(gin.Accounts)
	rtor := profls.GetEnumerator()
	for rtor.MoveNext() {
		dta := rtor.Current().Data().(Profile)
		prefx := strings.ToLower(dta.Title)
		for _, v := range dta.Clients {
			result[prefx + v.Name] = v.Secret
		}
	}

	return result, nil
}

func (cs *Clients) GetByID(id string) (oauth2.ClientInfo, error) {
	parts := strings.Split(id, ".")

	if len(parts) != 2 {
		return nil, errors.New("id doesn't have seperator")
	}

	rec, err := ctx.Profiles.FindFirst(byTitleID(parts[0]))

	if err != nil {
		return nil, err
	}

	prof := rec.Data().(Profile)

	for _, v := range prof.Clients {
		if v.Name == parts[1] {
			return v, nil
		}
	}

	return nil, errors.New("invalid client id")
}

func (cs *Clients) Set(id string, cli oauth2.ClientInfo) error {
	parts := strings.Split(id, ".")

	if len(parts) != 2 {
		return errors.New("id doesn't have seperator")
	}

	rec, err := ctx.Profiles.FindFirst(byTitleID(parts[0]))

	if err != nil {
		return err
	}

	prof := rec.Data().(Profile)

	for _, v := range prof.Clients {
		if v.Name == parts[1]{
			v = cli.(Client)
			break
		}
	}

	err = rec.Set(prof)

	if err != nil {
		return err
	}

	err = ctx.Profiles.Update(rec)

	if err != nil {
		return err
	}

	return ctx.Profiles.Save()
}
