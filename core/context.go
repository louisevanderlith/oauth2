package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Clients   husk.Tabler
	Users     husk.Tabler
	Forgotten husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Clients:   husk.NewTable(Client{}, serials.GobSerial{}),
		Users:     husk.NewTable(User{}, serials.GobSerial{}),
		Forgotten: husk.NewTable(Forgot{}, serials.GobSerial{}),
	}

	seed()
}

func Shutdown() {
	ctx.Users.Save()
	ctx.Clients.Save()
}

func seed() {
	err := ctx.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Users.Save()

	err = ctx.Clients.Seed("db/clients.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Clients.Save()
}
