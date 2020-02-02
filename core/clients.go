package core

import (
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
