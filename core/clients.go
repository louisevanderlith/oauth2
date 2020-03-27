package core

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/husk"
	"gopkg.in/oauth2.v3"
)

type Clients struct{}

func NewClientStore() oauth2.ClientStore {
	return &Clients{}
}

func GetAllClients() (husk.Collection, error) {
	return ctx.Clients.Find(1, 10, husk.Everything())
}

func GetClientAccounts() (gin.Accounts, error) {
	clnts, err := GetAllClients()

	if err != nil {
		return nil, err
	}

	result := make(gin.Accounts)

	rtor := clnts.GetEnumerator()
	for rtor.MoveNext() {
		dta := rtor.Current().Data().(Client)
		result[dta.ID] = dta.Secret
	}

	return result, nil
}

func (cs *Clients) GetByID(id string) (oauth2.ClientInfo, error) {
	rec, err := ctx.Clients.FindFirst(byID(id))

	if err != nil {
		return nil, err
	}

	return rec.Data().(Client), nil
}

func (cs *Clients) Set(id string, cli oauth2.ClientInfo) error {
	rec, err := ctx.Clients.FindFirst(byID(id))

	if err != nil {
		return err
	}

	err = rec.Set(cli.(Client))

	if err != nil {
		return err
	}

	err = ctx.Clients.Update(rec)

	if err != nil {
		return err
	}

	return ctx.Clients.Save()
}
